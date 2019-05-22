package testtree

import (
	"github.com/vektah/gqlparser/ast"
	"gotest.tools/assert"
	"testing"
)

func TestNewFile(t *testing.T) {
	before := &ast.Directive{
		Name: BeforeDirective,
	}
	after := &ast.Directive{
		Name: AfterDirective,
	}
	beforeAndAfterOp := &ast.OperationDefinition{
		Directives: []*ast.Directive{
			before,
			after,
		},
	}
	beforeOp := &ast.OperationDefinition{
		Directives: []*ast.Directive{
			before,
		},
	}
	afterOp := &ast.OperationDefinition{
		Directives: []*ast.Directive{
			after,
		},
	}
	testOp := &ast.OperationDefinition{}

	doc := &ast.QueryDocument{
		Operations: []*ast.OperationDefinition{
			beforeAndAfterOp,
			beforeOp,
			afterOp,
			testOp,
		},
	}

	name := "someFile.graphql"

	file := NewFile(doc, name)

	assert.Equal(t, file.name, name)
	assert.Equal(t, file.before[0], beforeAndAfterOp)
	assert.Equal(t, file.after[0], beforeAndAfterOp)
	assert.Equal(t, file.before[1], beforeOp)
	assert.Equal(t, file.after[1], afterOp)
	assert.Equal(t, file.tests[0], testOp)
}
