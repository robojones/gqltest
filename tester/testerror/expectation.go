package testerror

import (
	"fmt"
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
		joinPath(e.path),
	) + fmt.Sprintf(
		"Result from API\n%s",
		stringifyObject(e.result),
	)
}

func joinPath(path []string) string {
	p := ""
	for _, el := range path {
		p += "." + el
	}
	return p
}
