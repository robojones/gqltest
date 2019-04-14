package reader

import (
	"github.com/robojones/gqltest/test_util/tempdir"
	"gotest.tools/assert"
	"testing"
)

func TestReader_ReadDir(t *testing.T) {
	var (
		files = []string{"a.txt", "b.json", "c.graphql"}
		dirs  = []string{"d", "e", "f"}
	)

	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	// add contents to the temp dir
	for i, file := range files {
		files[i] = tempdir.File(t, dir, file, "")
	}
	for i, subdir := range dirs {
		dirs[i] = tempdir.Dir(t, dir, subdir)
	}

	// start testing
	r := NewReader()

	d := r.ReadDir(dir)
	assert.Equal(t, d.Name, dir)
	assert.DeepEqual(t, d.Files, files)
	assert.DeepEqual(t, d.Subdirs, dirs)
}
