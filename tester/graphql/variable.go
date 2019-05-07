package graphql

import (
	"github.com/vektah/gqlparser/ast"
	"strings"
)

func serializeVariableList(vars ast.VariableDefinitionList) string {
	if len(vars) == 0 {
		return ""
	}

	r := make([]string, len(vars))[:0]
	for _, v := range vars {
		r = append(r, serializeVariable(v))
	}

	return braceL + strings.Join(r, comma) + braceR
}

func serializeVariable(v *ast.VariableDefinition) string {
	r := dollar + v.Variable + colon + v.Type.String()

	if v.DefaultValue != nil {
		r += equal + v.DefaultValue.String()
	}

	return r
}
