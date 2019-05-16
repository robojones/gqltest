package testerror

import (
	"fmt"
	"github.com/robojones/gqltest/graphql/json"
)

type ComparisonError struct {
	actual   interface{}
	expected interface{}
}

func NewComparisonError(actual interface{}, expected interface{}) *ComparisonError {
	return &ComparisonError{
		expected: expected,
		actual:   actual,
	}
}

func (e *ComparisonError) Error() string {
	exp := json.StringifyValue(e.expected)
	act := json.StringifyValue(e.actual)

	return fmt.Sprintf("Expected %#v but instead got %#v", exp, act)
}
