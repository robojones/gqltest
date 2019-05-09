package json

import (
	"gotest.tools/assert"
	"testing"
)

func TestJsonValue(t *testing.T) {
	v := Value(t, "true").(bool)
	assert.Assert(t, v, true)
}
