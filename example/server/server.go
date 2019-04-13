package server

import (
	"github.com/caarlos0/env"
	"net"
	"net/http"
	"time"
)

type Config struct {
	Port string `env:"PORT" envDefault:"4000"`
}

type Server struct {
	closed bool
	listener net.Listener
	http *http.Server
}

func (s *Server) Serve() error {
	return s.http.Serve(s.listener.(*net.TCPListener))
}

func (s *Server) Listen() error {
	if s.closed {
		return http.ErrServerClosed
	}

	addr := s.http.Addr

	ln, err := net.Listen("tcp", addr)
	s.listener = ln

	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Close() error {
	return s.http.Close()
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := env.Parse(config)
	return config, err
}

func NewServer(mux *http.ServeMux, config *Config) *Server {
	return &Server{
		http: &http.Server{
			Addr:           ":" + config.Port,
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}
