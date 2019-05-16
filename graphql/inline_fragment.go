package graphql

import (
	"github.com/vektah/gqlparser/ast"
)

func serializeInlineFragment(frag *ast.InlineFragment, indent int) string {
	r := spread

	if frag.TypeCondition != "" {
		r += on + frag.TypeCondition
	}

	return r +
		serializeDirectiveList(frag.Directives) +
		serializeSelectionSet(frag.SelectionSet, indent)
}
