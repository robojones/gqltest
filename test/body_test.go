package test

import (
	"github.com/vektah/gqlparser/ast"
	"gotest.tools/assert"
	"testing"
)

func TestTest_Body(t *testing.T) {
	doc := getTestDoc()
	selection := doc.Operations[0].SelectionSet[0]
	assert.Equal(t, len(selection.(*ast.Field).Directives), 1)

	test := NewTest(doc)

	orig := test.Body()
	assert.Assert(t, test.Body() == orig)

	assert.NilError(t, test.ParseExpectations())
	assert.Assert(t, test.Body() != orig)
}
