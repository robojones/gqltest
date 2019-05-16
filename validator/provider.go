package validator

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewValidator,
)
