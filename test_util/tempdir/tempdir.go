package tempdir

import (
	"gotest.tools/assert"
	"io/ioutil"
	"os"
	"testing"
)

// Create a new directory and return its path.
func Create(t *testing.T) string {
	t.Helper()

	dir, err := ioutil.TempDir(os.TempDir(), "gqltest_")
	assert.NilError(t, err)
	return dir
}

// Remove a directory and its contents.
func Remove(t *testing.T, dir string) {
	t.Helper()

	err := os.RemoveAll(dir)
	assert.NilError(t, err)
}
