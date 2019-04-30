package json

import (
	"github.com/robojones/gqltest/test_util/json"
	"gotest.tools/assert"
	"testing"
)

func TestDetectType(t *testing.T) {
	assert.Equal(t, DetectType(json.Value("true")), "Boolean")
	assert.Equal(t, DetectType(json.Value("10")), "Number")
	assert.Equal(t, DetectType(json.Value("5.5")), "Number")
	assert.Equal(t, DetectType(json.Value("{}")), "Object")
	assert.Equal(t, DetectType(json.Value("[]")), "Array")
	assert.Equal(t, DetectType(json.Value("null")), "null")
}
