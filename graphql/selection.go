package graphql

import (
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/ast"
	"strings"
)

func serializeSelectionSet(sels ast.SelectionSet, indent int) string {
	if len(sels) == 0 {
		return ""
	}

	indent += 1

	r := make([]string, len(sels))[:0]
	for _, sel := range sels {
		r = append(r, strings.Repeat(tab, indent)+serializeSelection(sel, indent))
	}

	indent -= 1

	return bracketL +
		newline +
		strings.Join(r, newline) +
		newline +
		strings.Repeat(tab, indent) +
		bracketR
}

func serializeSelection(sel ast.Selection, indent int) string {
	switch sel.(type) {
	case *ast.Field:
		return serializeField(sel.(*ast.Field), indent)
	case *ast.InlineFragment:
		return serializeInlineFragment(sel.(*ast.InlineFragment), indent)
	case *ast.FragmentSpread:
		return serializeFragmentSpread(sel.(*ast.FragmentSpread))
	default:
		panic(errors.Errorf("Unknown selection %#v", sel))
	}
}
