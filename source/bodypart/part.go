package bodypart

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/source"
)

type Part struct {
	kind   string
	Source *source.Source
	node   ast.Node
}

func NewPart(node ast.Node, s *source.Source) *Part {
	return &Part{
		kind:   node.GetKind(),
		Source: s,
		node:   node,
	}
}

func (p *Part) Content() []byte {
	l := p.node.GetLoc()

	return p.Source.Body[l.Start:l.End]
}
