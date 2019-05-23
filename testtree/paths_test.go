package testtree

import (
	"github.com/vektah/gqlparser/ast"
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func includes(paths []Path, name string) bool {
	path := NewPath(name)
	path = path[:len(path) - 1]

	for _, p := range paths {
		if reflect.DeepEqual(p, path) {
			return true
		}
	}
	return false
}

func TestGetTestPaths(t *testing.T) {
	f := &File{
		tests: []*ast.OperationDefinition{
			{},
		},
	}
	name := "path/to/coolness.graphql"
	name2 := "path/to/wisdom.graphql"

	tree := NewTree()
	tree.SetFile(name, f)
	tree.SetFile(name2, f)

	paths := GetTestPaths(tree)

	assert.Assert(t, includes(paths, name))
	assert.Assert(t, includes(paths, name2))
}
