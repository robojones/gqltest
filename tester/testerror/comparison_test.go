package testerror

import (
	"gotest.tools/assert"
	"testing"
)

func TestNewComparisonError(t *testing.T) {
	const actual = 10
	const expected = 234

	err := NewComparisonError(actual, expected)

	assert.ErrorType(t, err, new(ComparisonError))
	assert.Equal(t, err.actual, actual)
	assert.Equal(t, err.expected, expected)
}

func TestComparisonError_Error(t *testing.T) {
	const actual = "halp"
	const expected = "fooo"

	err := NewComparisonError(actual, expected)

	assert.ErrorContains(t, err, actual)
	assert.ErrorContains(t, err, expected)
}
