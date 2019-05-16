package config

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewConfig,
)
