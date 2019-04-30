package testerror

import (
	"fmt"
	"github.com/robojones/gqltest/tester/json"
)

type TypeError struct {
	actualType   string
	expectedType string
}

func NewTypeError(actualValue interface{}, expectedType string) *TypeError {
	return &TypeError{actualType: json.DetectType(actualValue), expectedType: expectedType}
}

func (e *TypeError) Error() string {
	return fmt.Sprintf(
		"Expected type %s but instead got %s",
		e.expectedType,
		e.actualType,
	)
}
