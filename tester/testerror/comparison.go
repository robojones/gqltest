package testerror

import (
	"fmt"
)

const Prefix = ""
const Indent = "  "

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
	exp := stringifyValue(e.expected)
	act := stringifyValue(e.actual)

	return fmt.Sprintf("Expected %#v but instead got %#v", exp, act)
}
