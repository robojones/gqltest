//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/robojones/gqltest/config"
	"github.com/robojones/gqltest/env"
	"github.com/robojones/gqltest/example/api"
	"github.com/robojones/gqltest/example/server"
	"github.com/robojones/gqltest/tester"
	"github.com/robojones/gqltest/validator"
)

func InitTester(testdir config.WorkinDirectoryName) (*tester.Tester, error) {
	wire.Build(
		config.Provider,
		env.Provider,
		tester.Provider,
		validator.Provider,

		// config bindings
		wire.Bind(new(validator.ValidatorConfig), new(config.Config)),
		wire.Bind(new(tester.TesterConfig), new(config.Config)),

		// env bindings
		wire.Bind(new(validator.ValidatorEnv), new(env.Env)),
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
