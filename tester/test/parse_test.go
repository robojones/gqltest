package test

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/expectation"
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/parser"
	"gotest.tools/assert"
	"testing"
)

var source = &ast.Source{
	Name:  "testSource.graphql",
	Input: "query { hello @expectBoolean(v: true) }",
}

func getTestDoc() *ast.QueryDocument {
	doc, err := parser.ParseQuery(source)

	if err != nil {
		panic(errors.WithStack(err))
	}

	return doc
}

func TestTest_ParseExpectations(t *testing.T) {
	doc := getTestDoc()
	op := doc.Operations[0]

	test := NewTest(doc)

	assert.NilError(t, test.ParseExpectations())
	assert.Assert(t, test.Expectations[0].(*expectation.ValueExpectation) != nil)
	assert.Equal(t, len(op.SelectionSet[0].(*ast.Field).Directives), 0)
}
