package bodypart

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/source"
	"gotest.tools/assert"
	"testing"
)

const (
	query    = "query {foo}"
	mutation = "mutation {bar}"
	kind     = "TestNode"
)

var (
	testSource = &source.Source{
		Body: []byte(query + mutation),
	}
	queryNode ast.Node = &astNodeMock{
		Start: 0,
		End:   len(query),
	}
	mutationNode ast.Node = &astNodeMock{
		Start: len(query),
		End:   len(query) + len(mutation),
	}
)

type astNodeMock struct {
	Start int
	End   int
}

func (n *astNodeMock) GetKind() string {
	return kind
}

func (n *astNodeMock) GetLoc() *ast.Location {
	return ast.NewLocation(&ast.Location{
		Start:  n.Start,
		End:    n.End,
		Source: testSource,
	})
}

func TestNewPart(t *testing.T) {
	part := NewPart(queryNode, testSource)

	assert.Equal(t, part.kind, kind)
	assert.Equal(t, part.Source, testSource)
	assert.Equal(t, part.node, queryNode)
}

func TestPart_Content(t *testing.T) {
	queryPart := NewPart(queryNode, testSource)
	assert.Equal(t, string(queryPart.Content()), query)

	mutationPart := NewPart(mutationNode, testSource)
	assert.Equal(t, string(mutationPart.Content()), mutation)
}
