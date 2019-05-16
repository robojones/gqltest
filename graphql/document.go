package graphql

import (
	"github.com/vektah/gqlparser/ast"
	"strings"
)

const (
	tab     = "  "
	space   = " "
	newline = "\n"

	braceL   = "("
	braceR   = ")"
	bracketL = " {"
	bracketR = "}"

	equal = " = "
	colon = ": "
	comma = ", "

	spread = "..."
	dollar = "$"
	at     = " @"

	on       = " on "
	fragment = "fragment "
)

func SerializeDocument(doc *ast.QueryDocument) string {
	r := make([]string, len(doc.Operations))[:0]

	for _, frag := range doc.Fragments {
		r = append(r, serializeFragment(frag, 0))
	}

	for _, op := range doc.Operations {
		r = append(r, serializeOperation(op, 0))
	}

	return strings.Join(r, newline)
}
