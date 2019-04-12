package tempfile

import (
	"gotest.tools/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

// Create a temporary file in the provided testing directory with the provided name.
func Create(t *testing.T, testDir string, name string, body string) (filepath string) {
	p := path.Join(testDir, name)
	err := ioutil.WriteFile(p, []byte(body), 0600)
	assert.NilError(t, err)
	return p
}

// Removes a file.
func Remove(t *testing.T, filepath string) {
	err := os.Remove(filepath)
	assert.NilError(t, err)
}
