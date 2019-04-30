package testerror

import (
	"fmt"
	"github.com/vektah/gqlparser/ast"
)

type TestError struct {
	err       error
	operation *ast.OperationDefinition
	directive *ast.Directive
}

func NewTestError(err error, operation *ast.OperationDefinition, directive *ast.Directive) *TestError {
	return &TestError{err: err, operation: operation, directive: directive}
}

func (e *TestError) Error() string {
	return fmt.Sprintf(
		"%s\n",
		e.err.Error(),
	) + fmt.Sprintf(
		"Expected in %s (line %d, column %d)\n%s",
		e.directive.Position.Src.Name,
		e.directive.Position.Line,
		e.directive.Position.Column,
		highlight(e.directive.Position, e.operation.Position),
	)
}
