package expectation

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/request"
	"github.com/robojones/gqltest/tester/testerror"
	"reflect"
)

type ValueExpectation struct {
	path  []string
	value *interface{}
}

func NewValueExpectation(
	path []string,
	value *interface{},
) Expectation {
	if len(path) < 1 {
		panic(errors.New("Path for expectation must contain at least one element."))
	}

	return &ValueExpectation{
		path:  path,
		value: value,
	}
}

func (exp *ValueExpectation) Check(result *request.Result) error {
	pathRemaining := exp.path
	scope := reflect.ValueOf(*result)

	for len(pathRemaining) > 0 {
		key := reflect.ValueOf(pathRemaining[0])
		pathRemaining = pathRemaining[1:]

		if s, ok := scope.Interface().(map[string]interface{}); ok {
			scope = reflect.ValueOf(s)
		} else {
			val := scope.Interface()
			typeErr := testerror.NewTypeError(&val, "Object")
			pathTaken := exp.path[:len(exp.path)-len(pathRemaining)]
			return testerror.NewExpectationError(typeErr, pathTaken, result)
		}

		scope = scope.MapIndex(key)
	}

	actual := scope.Interface()

	if !reflect.DeepEqual(actual, *exp.value) {
		compErr := testerror.NewComparisonError(*exp.value, actual)
		return testerror.NewExpectationError(compErr, exp.path, result)
	}

	return nil
}
