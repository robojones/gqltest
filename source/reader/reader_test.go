package reader

import (
	"gotest.tools/assert"
	"testing"
)

func TestNewReader(t *testing.T) {
	reader := NewReader()

	assert.Assert(t, reader != nil)
}
