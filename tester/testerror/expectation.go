package testerror

import (
	"fmt"
	"github.com/robojones/gqltest/graphql/json"
	"strings"
)

type ExpectationError struct {
	err    error
	result interface{}
	path   []string
}

func NewExpectationError(
	err error,
	path []string,
	result interface{},
) *ExpectationError {
	return &ExpectationError{
		err:    err,
		path:   path,
		result: result,
	}
}

func (e *ExpectationError) Cause() error {
	return e.err
}

func (e *ExpectationError) Error() string {
	return fmt.Sprintf(
		"%s\n",
		e.err.Error(),
	) + fmt.Sprintf(
		"Field: %s\n\n",
		strings.Join(e.path, "."),
	) + fmt.Sprintf(
		"Result from API\n%s",
		json.StringifyObject(e.result),
	)
}
