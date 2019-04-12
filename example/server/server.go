package server

import (
	"github.com/caarlos0/env"
	"log"
	"net/http"
)

type Config struct {
	Port string `env:"PORT" envDefault:"4000"`
}

type Server struct {
	Mux    *http.ServeMux
	Config *Config
}

func (s *Server) Listen() error {

	log.Printf("listening on http://localhost:%s", s.Config.Port)

	addr := ":" + s.Config.Port
	return http.ListenAndServe(addr, s.Mux)
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := env.Parse(config)
	return config, err
}

func NewServer(mux *http.ServeMux, config *Config) *Server {
	return &Server{
		Mux:    mux,
		Config: config,
	}
}
