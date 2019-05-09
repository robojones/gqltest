package graphql

import "github.com/vektah/gqlparser/ast"

func serializeField(field *ast.Field, indent int) string {
	r := ""

	if field.Alias != field.Name {
		r = field.Alias + colon
	}

	return r +
		field.Name +
		serializeArgumentList(field.Arguments) +
		serializeDirectiveList(field.Directives) +
		serializeSelectionSet(field.SelectionSet, indent)
}
