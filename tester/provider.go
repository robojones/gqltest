package tester

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewTester,
)
