package testerror

import (
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/parser"
	"gotest.tools/assert"
	"testing"
)

const (
	exampleDirective = "@expect(v: 5)"
	exampleQuery     = "query Test {\n" +
		"  wazzup " + exampleDirective + "\n" +
		"}"
)

var (
	testSource = &ast.Source{
		Name:  "testfile.graphql",
		Input: "\n" + exampleQuery + "\n",
	}
	testDocument  *ast.QueryDocument
	testOperation *ast.OperationDefinition
	testDirective *ast.Directive
)

func init() {
	doc, err := parser.ParseQuery(testSource)

	if err != nil {
		panic(errors.WithStack(err))
	}

	testDocument = doc
	testOperation = testDocument.Operations[0]
	field := testOperation.SelectionSet[0].(*ast.Field)
	testDirective = field.Directives[0]

}

func TestNewTestError(t *testing.T) {
	e := NewTestError(orig, testOperation, testDirective)

	assert.ErrorType(t, e, new(TestError))
	assert.Equal(t, e.err, orig)
	assert.Equal(t, e.operation, testOperation)
	assert.Equal(t, e.directive, testDirective)
}

func TestTestError_Error(t *testing.T) {
	e := NewTestError(orig, testOperation, testDirective)

	assert.ErrorContains(t, e, highlight(testDirective))
	assert.ErrorContains(t, e, orig.Error())
	assert.ErrorContains(t, e, testOperation.Position.Src.Name)
}
