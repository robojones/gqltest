package testerror

import (
	"gotest.tools/assert"
	"testing"
)

func TestNewTypeError(t *testing.T) {
	const (
		expected = "object"
		actual   = "boolean"
	)

	e := NewTypeError(actual, expected)

	assert.ErrorType(t, e, new(TypeError))
	assert.Equal(t, e.actualType, actual)
	assert.Equal(t, e.expectedType, expected)
}

func TestTypeError_Error(t *testing.T) {
	const (
		expected = "object"
		actual   = "boolean"
	)

	e := NewTypeError(actual, expected)

	assert.ErrorContains(t, e, expected)
	assert.ErrorContains(t, e, actual)

}
