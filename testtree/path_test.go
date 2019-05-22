package testtree

import (
	"gotest.tools/assert"
	"testing"
)

func testNewPath(t *testing.T, inputPath string, expect Path) {
	t.Helper()

	r := NewPath(inputPath)
	assert.DeepEqual(t, r, expect)
}

func TestNewPath(t *testing.T) {
	testNewPath(t, "some/path/file.jpg", Path{"some", "path", "file.jpg"})
	testNewPath(t, "/some/path/file.jpg", Path{"some", "path", "file.jpg"})
	testNewPath(t, "/", Path{})
	testNewPath(t, ".", Path{})
	testNewPath(t, "", Path{})
}
