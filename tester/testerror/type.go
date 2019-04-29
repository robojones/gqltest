package testerror

import (
	"fmt"
)

type TypeError struct {
	actualType   string
	expectedType string
}

func NewTypeError(actualType string, expectedType string) *TypeError {
	return &TypeError{actualType: actualType, expectedType: expectedType}
}

func (e *TypeError) Error() string {
	return fmt.Sprintf(
		"Expected type %s but instead got %s",
		e.expectedType,
		e.actualType,
	)
}
