package expectation

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/request"
	"github.com/robojones/gqltest/tester/testerror"
	"reflect"
)

func ExpectValue(
	path []string,
	expected interface{},
) Expectation {
	if len(path) < 1 {
		panic(errors.New("Path for expectation must contain at least one element."))
	}

	return func(result *request.Result) error {
		pathRemaining := path
		scope := reflect.ValueOf(*result)

		for len(pathRemaining) > 0 {
			key := reflect.ValueOf(pathRemaining[0])
			pathRemaining = pathRemaining[1:]

			if s, ok := scope.Interface().(map[string]interface{}); ok {
				scope = reflect.ValueOf(s)
			} else {
				val := scope.Interface()
				typeErr := testerror.NewTypeError(&val, "Object")
				pathTaken := path[:len(path)-len(pathRemaining)]
				return testerror.NewExpectationError(typeErr, pathTaken, result)
			}

			scope = scope.MapIndex(key)
		}

		actual := scope.Interface()

		if !reflect.DeepEqual(actual, expected) {
			compErr := testerror.NewComparisonError(expected, actual)
			return testerror.NewExpectationError(compErr, path, result)
		}

		return nil
	}
}
