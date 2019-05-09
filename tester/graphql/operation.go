package graphql

import (
	"github.com/vektah/gqlparser/ast"
)

func serializeOperation(op *ast.OperationDefinition, indent int) string {
	r := string(op.Operation)

	if op.Name != "" {
		r += space + op.Name
	}

	return r +
		serializeVariableList(op.VariableDefinitions) +
		serializeDirectiveList(op.Directives) +
		serializeSelectionSet(op.SelectionSet, indent)
}
