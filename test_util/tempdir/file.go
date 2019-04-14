package tempdir

import (
	"gotest.tools/assert"
	"io/ioutil"
	"path"
	"testing"
)

const perm = 0600

// Create a temporary file in the provided testing directory with the provided name.
func File(t *testing.T, testDir string, name string, body string) (filepath string) {
	p := path.Join(testDir, name)
	err := ioutil.WriteFile(p, []byte(body), perm)
	assert.NilError(t, err)
	return p
}
