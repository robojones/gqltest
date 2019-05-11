//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/robojones/gqltest/config"
	"github.com/robojones/gqltest/example/api"
	"github.com/robojones/gqltest/example/server"
	"github.com/robojones/gqltest/source/reader"
	"github.com/robojones/gqltest/tester"
)

func InitTester(testdir config.WorkinDirectoryName) (*tester.Tester, error) {
	wire.Build(
		config.Provider,
		tester.Provider,
		reader.Provider,
	)

	return nil, nil
}

func InitTestServer() *server.Server {
	wire.Build(
		server.Provider,
		api.Provider,
	)

	return nil
}
