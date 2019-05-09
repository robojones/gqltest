package test

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/expectation"
	"github.com/robojones/gqltest/tester/testerror"
	"github.com/vektah/gqlparser/ast"
)

type Test struct {
	// GraphQL AST for the test.
	Document *ast.QueryDocument

	// Expectations for the request result.
	Expectations []expectation.Expectation

	// Errors returned by the expectations.
	Errors  []*testerror.TestError
	Success bool
}

func NewTest(doc *ast.QueryDocument) *Test {
	if doc == nil {
		panic(errors.New("doc must not be nil"))
	}

	if len(doc.Operations) != 1 {
		panic(errors.New("test must contain exactly one operation"))
	}

	return &Test{
		Document: doc,
	}
}
