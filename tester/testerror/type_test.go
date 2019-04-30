package testerror

import (
	"fmt"
	"github.com/robojones/gqltest/test_util/json"
	"gotest.tools/assert"
	"testing"
)

func TestNewTypeError(t *testing.T) {
	const (
		expected = "Object"
		actual   = "Boolean"
	)
	var value = json.Value("true")

	e := NewTypeError(value, expected)

	assert.ErrorType(t, e, new(TypeError))
	assert.Equal(t, e.actualType, actual)
	assert.Equal(t, e.expectedType, expected)
}

func TestTypeError_Error(t *testing.T) {
	const (
		expected = "Object"
		actual   = "Boolean"
	)
	var value = json.Value("true")

	e := NewTypeError(value, expected)

	assert.Error(t, e, fmt.Sprintf("Expected type %s but instead got %s", expected, actual))
}
