package reader

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewReader,
)
