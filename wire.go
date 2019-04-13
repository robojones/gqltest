//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/robojones/gqltest/config"
	"github.com/robojones/gqltest/source/reader"
	"github.com/robojones/gqltest/tester"
)

func InitTester(testdir config.WD) (*tester.Tester, error) {
	wire.Build(
		config.Provider,
		tester.Provider,
		reader.Provider,
	)

	return nil, nil
}
