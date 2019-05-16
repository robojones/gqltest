package test

import (
	"github.com/vektah/gqlparser/ast"
)

func (t *Test) Operation() *ast.OperationDefinition {
	return t.Document.Operations[0]
}
