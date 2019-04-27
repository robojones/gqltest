package reader

import (
	"github.com/robojones/gqltest/test_util/tempdir"
	"gotest.tools/assert"
	"testing"
)

func TestReader_Read(t *testing.T) {
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

	reader := NewReader()

	r := reader.Read(dir)

	for _, file := range files {
		_, ok := r[file]
		assert.Assert(t, ok, "file %s not in map")
	}
}
