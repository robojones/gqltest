package reader

import (
	"github.com/robojones/gqltest/test_util/tempdir"
	"gotest.tools/assert"
	"path"
	"testing"
)

func TestReader_Read(t *testing.T) {
	var (
		files    = []string{"a.txt", "b.graphql", "c.graphql"}
		matching = files[1:]
		dirs     = []string{"d"}
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

	r := reader.Read(path.Join(dir, "*.graphql"))

	for _, file := range matching {
		_, ok := r[file]
		assert.Assert(t, ok, "file %s not in map", file)
	}

	for _, dir := range dirs {
		_, ok := r[dir]
		assert.Assert(t, !ok, "dir %s in map, but it shouldn' be there.", dir)
	}
}
