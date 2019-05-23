package testtree

import (
	"github.com/vektah/gqlparser/ast"
	"gotest.tools/assert"
	"testing"
)

func TestAddChild(t *testing.T) {
	parent := NewNode()
	name := "childName"

	AddChild(parent, name)

	assert.Assert(t, parent.children[name] != nil)
}

func TestNewNode(t *testing.T) {
	node := NewNode()
	assert.Assert(t, node != nil)
}

func TestNode_IsTestNode(t *testing.T) {
	node := NewNode()

	node.Files["somefile"] = &File{}
	assert.Assert(t, !node.IsTestNode())

	node.Files["otherfile"] = &File{
		tests: []*ast.OperationDefinition{{}},
	}
	assert.Assert(t, node.IsTestNode())
}
