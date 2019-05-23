package testtree

import (
	"gotest.tools/assert"
	"testing"
)

func TestNewNodes(t *testing.T) {
	tree := NewTree()
	p := Path{"test", "node", "path"}
	n := NewNodes(tree, p)

	assert.Equal(t, n.index, 0)
	assert.Equal(t, n.current, tree.root)
	assert.DeepEqual(t, n.path, p)
}

func TestNodes_Value(t *testing.T) {
	tree := NewTree()
	nodes := NewNodes(tree, Path{})

	assert.Equal(t, nodes.Value(), nodes.current)
}

func TestNodes_Next(t *testing.T) {
	tree := NewTree()
	filename := "file.graphql"
	file := &File{}
	tree.SetFile("some/cool/" + filename, file)
	p := Path{"some", "cool"}

	nodes := NewNodes(tree, p)

	n := 0
	for nodes.Next() {
		n += 1
		assert.Equal(t, nodes.index, n)
	}
	node := nodes.current

	assert.Equal(t, nodes.index, len(p))
	assert.Equal(t, node.Files[filename], file)
}

func TestNodes_Previous(t *testing.T) {
	tree := NewTree()
	filename := "file.graphql"
	file := &File{}
	tree.SetFile("some/cool/" + filename, file)
	p := Path{"some", "cool"}

	nodes := NewNodes(tree, p)
	n := 0
	for nodes.Next() {
		n += 1
	}
	assert.Equal(t, n, len(p))
	assert.Assert(t, nodes.current != tree.root)

	n2 := 0
	for nodes.Previous() {
		n2 += 1
	}
	assert.Equal(t, n2, len(p))
	assert.Equal(t, nodes.current, tree.root)
}
