package testtree

import (
	"gotest.tools/assert"
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := NewTree()
	assert.Assert(t, tree.root != nil)
}

func TestTree_SetFile(t *testing.T) {
	p := "path/to/file.graphql"
	f := &File{}
	tree := NewTree()

	tree.SetFile(p, f)

	r := tree.root.children["path"].children["to"].files["file.graphql"]
	assert.Equal(t, r, f)
}
