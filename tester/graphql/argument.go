package graphql

import (
	"github.com/vektah/gqlparser/ast"
	"strings"
)

func serializeArgumentList(args ast.ArgumentList) string {
	if len(args) == 0 {
		return ""
	}

	r := make([]string, len(args))[:0]
	for _, arg := range args {
		r = append(r, arg.Name+colon+arg.Value.String())
	}

	return braceL + strings.Join(r, comma) + braceR
}
