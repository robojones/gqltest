package server

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewServer,
	NewConfig,
	NewServeMux,
)
