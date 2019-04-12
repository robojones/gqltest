//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/robojones/gqltest/example/api"
	"github.com/robojones/gqltest/example/server"
)

func InitServer() (*server.Server, error) {
	wire.Build(
		server.Provider,
		api.Provider,
	)

	return nil, nil
}
