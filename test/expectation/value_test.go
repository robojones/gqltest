package expectation

import (
	"github.com/robojones/gqltest/tester/request"
	"github.com/robojones/gqltest/tester/testerror"
	"github.com/vektah/gqlparser/ast"
	"gotest.tools/assert"
	"testing"
)

func TestNewValueExpectation(t *testing.T) {
	directive := &ast.Directive{}
	path := []string{"data", "foo", "bar"}
	var val = "cool"

	exp := NewValueExpectation(directive, path, val).(*ValueExpectation)

	assert.Equal(t, exp.directive, directive)
	assert.Equal(t, exp.value.(string), val)
	assert.DeepEqual(t, exp.path, path)
}

func TestValueExpectation_Check(t *testing.T) {
	directive := &ast.Directive{}
	path := []string{"data", "foo", "bar"}
	var val = "cool"

	result, err := request.ParseResult([]byte(`{ "data": { "foo": { "bar": "cool" }}}`))

	assert.NilError(t, err)

	exp := NewValueExpectation(directive, path, val)
	err = exp.Check(result)

	assert.NilError(t, err)
}

func TestValueExpectation_CheckExpectValueComparisonError(t *testing.T) {
	path := []string{"data", "foo", "bar"}
	var val interface{} = "cool"

	result, err := request.ParseResult([]byte(`{ "data": { "foo": { "bar": "not cool" }}}`))

	assert.NilError(t, err)

	exp := NewValueExpectation(nil, path, &val)
	err = exp.Check(result)

	assert.ErrorType(t, err, new(testerror.ExpectationError))
	assert.ErrorType(t, err.(*testerror.ExpectationError).Cause(), new(testerror.ComparisonError))
}

func TestValueExpectation_CheckExpectTypeError(t *testing.T) {
	path := []string{"data", "foo", "bar"}
	var val = "some value"

	result, err := request.ParseResult([]byte(`{ "data": true }`))

	assert.NilError(t, err)

	exp := NewValueExpectation(nil, path, val)
	err = exp.Check(result)

	assert.ErrorType(t, err, new(testerror.ExpectationError))
	assert.ErrorType(t, err.(*testerror.ExpectationError).Cause(), new(testerror.TypeError))
}

func TestValueExpectation_Directive(t *testing.T) {
	directive := &ast.Directive{}

	exp := NewValueExpectation(directive, []string{"asdf"}, "")

	assert.Equal(t, exp.Directive(), directive)
}
