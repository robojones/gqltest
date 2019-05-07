package graphql

import (
	"github.com/vektah/gqlparser/ast"
	"strings"
)

func serializeDirectiveList(dirs ast.DirectiveList) string {
	if len(dirs) == 0 {
		return ""
	}

	r := make([]string, len(dirs))[:0]

	for _, dir := range dirs {
		r = append(r, serializeDirective(dir))
	}

	return strings.Join(r, space)
}

func serializeDirective(dir *ast.Directive) string {
	return at + dir.Name + serializeArgumentList(dir.Arguments)
}
