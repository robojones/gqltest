package testerror

import (
	"fmt"
	"github.com/pkg/errors"
)

type TypeError struct {
	actualType   string
	expectedType string
}

func NewTypeError(actualValue *interface{}, expectedType string) *TypeError {
	return &TypeError{actualType: DetectType(actualValue), expectedType: expectedType}
}

func (e *TypeError) Error() string {
	return fmt.Sprintf(
		"Expected type %s but instead got %s",
		e.expectedType,
		e.actualType,
	)
}

func DetectType(v *interface{}) string {
	switch (*v).(type) {
	case string:
		return "String"
	case float64:
		return "Number"
	case map[string]interface{}:
		return "Object"
	case []interface{}:
		return "Array"
	case bool:
		return "Boolean"
	case nil:
		return "null"
	default:
		panic(errors.Errorf("Unknown type %T of value %#v", v, v))
	}
}
