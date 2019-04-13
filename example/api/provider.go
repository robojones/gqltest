package api

import (
	"github.com/google/wire"
	"github.com/robojones/gqltest/example/api/query"
)

var Provider = wire.NewSet(
	NewHandler,
	query.NewResolver,
)
