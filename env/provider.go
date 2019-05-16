package env

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewEnv,
)
