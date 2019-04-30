package json

import (
	"gotest.tools/assert"
	"testing"
)

var (
	exampleObject interface{} = map[string]interface{}{
		"foo": "bar",
	}
)

func TestStringifyValue(t *testing.T) {
	s := StringifyValue(exampleObject)

	assert.Equal(t, s, `{"foo":"bar"}`)
}

func TestStringifyObject(t *testing.T) {
	s := StringifyObject(exampleObject)

	assert.Equal(t, s, "{\n  \"foo\": \"bar\"\n}")
}
