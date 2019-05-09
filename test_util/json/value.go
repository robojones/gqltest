package json

import (
	"encoding/json"
	"gotest.tools/assert"
	"testing"
)

func Value(t *testing.T, s string) interface{} {
	t.Helper()

	v := new(interface{})
	assert.NilError(t, json.Unmarshal([]byte(s), v))

	return *v
}
