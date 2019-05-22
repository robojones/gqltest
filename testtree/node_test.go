package testtree

import (
	"gotest.tools/assert"
	"testing"
)

func TestAddChild(t *testing.T) {
	parent := NewNode()
	name := "childName"

	AddChild(parent, name)

	assert.Assert(t, parent.children[name] != nil)
}
