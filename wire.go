//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/robojones/gqltest/source/reader"
	"github.com/robojones/gqltest/tester"
)

func InitTester() *tester.Tester {
	wire.Build(
		tester.Provider,
		reader.Provider,
	)

	return nil
}
