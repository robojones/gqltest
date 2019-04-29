package expectation

import (
	"github.com/robojones/gqltest/tester/request"
	"github.com/robojones/gqltest/tester/testerror"
	"gotest.tools/assert"
	"testing"
)

func TestExpectValue(t *testing.T) {
	path := []string{"data", "foo", "bar"}
	val := "cool"

	result, err := request.ParseResult([]byte(`{ "data": { "foo": { "bar": "cool" }}}`))

	assert.NilError(t, err)

	e := ExpectValue(path, val)
	err = e(result)

	assert.NilError(t, err)
}

func TestExpectValueComparisonError(t *testing.T) {
	path := []string{"data", "foo", "bar"}
	val := "cool"

	result, err := request.ParseResult([]byte(`{ "data": { "foo": { "bar": "not cool" }}}`))

	assert.NilError(t, err)

	e := ExpectValue(path, val)
	err = e(result)

	assert.ErrorType(t, err, new(testerror.ExpectationError))
	assert.ErrorType(t, err.(*testerror.ExpectationError).Cause(), new(testerror.ComparisonError))
}

func TestExpectTypeError(t *testing.T) {
	path := []string{"data", "foo", "bar"}
	val := "cool"

	result, err := request.ParseResult([]byte(`{ "data": true }`))

	assert.NilError(t, err)

	e := ExpectValue(path, val)
	err = e(result)

	assert.ErrorType(t, err, new(testerror.ExpectationError))
	assert.ErrorType(t, err.(*testerror.ExpectationError).Cause(), new(testerror.TypeError))
}
