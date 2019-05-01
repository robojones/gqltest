package test

import (
	"github.com/robojones/gqltest/tester/expectation"
	"github.com/robojones/gqltest/tester/testerror"
	"github.com/vektah/gqlparser/ast"
)

type Test struct {
	Body string

	// Operation for the test.
	Operation *ast.OperationDefinition

	Directives []*ast.Position

	// Expectations for the request result.
	Expectations []expectation.Expectation

	// Errors returned by the expectations.
	Errors  []*testerror.TestError
	Success bool
}

func NewTest(operation *ast.OperationDefinition, body string) *Test {
	return &Test{
		Body:      body,
		Operation: operation,
	}
}
