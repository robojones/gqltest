package json

import (
	"github.com/robojones/gqltest/test_util/json"
	"gotest.tools/assert"
	"testing"
)

func TestDetectType(t *testing.T) {
	assert.Equal(t, DetectType(json.Value(t, "true")), "Boolean")
	assert.Equal(t, DetectType(json.Value(t, "10")), "Number")
	assert.Equal(t, DetectType(json.Value(t, "5.5")), "Number")
	assert.Equal(t, DetectType(json.Value(t, "{}")), "Object")
	assert.Equal(t, DetectType(json.Value(t, "[]")), "Array")
	assert.Equal(t, DetectType(json.Value(t, "null")), "null")
}
