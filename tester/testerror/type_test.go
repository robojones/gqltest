package testerror

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"gotest.tools/assert"
	"testing"
)

func jsonValue(s string) *interface{} {
	v := new(interface{})
	if err := json.Unmarshal([]byte(s), v); err != nil {
		panic(errors.WithStack(err))
	}
	return v
}

func TestDetectType(t *testing.T) {
	assert.Equal(t, DetectType(jsonValue("true")), "Boolean")
	assert.Equal(t, DetectType(jsonValue("10")), "Number")
	assert.Equal(t, DetectType(jsonValue("5.5")), "Number")
	assert.Equal(t, DetectType(jsonValue("{}")), "Object")
	assert.Equal(t, DetectType(jsonValue("[]")), "Array")
	assert.Equal(t, DetectType(jsonValue("null")), "null")
}

func TestNewTypeError(t *testing.T) {
	const (
		expected = "Object"
		actual   = "Boolean"
	)
	var value = jsonValue("true")

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
	var value = jsonValue("true")

	e := NewTypeError(value, expected)

	assert.Error(t, e, fmt.Sprintf("Expected type %s but instead got %s", expected, actual))
}
