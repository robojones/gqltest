package server

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/graph-gophers/graphql-go/relay"
	"net/http"
)

func NewServeMux(api *relay.Handler) *http.ServeMux {
	mux := &http.ServeMux{}

	mux.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	mux.Handle("/query", api)

	return mux
}
