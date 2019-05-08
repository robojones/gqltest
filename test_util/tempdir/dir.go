package tempdir

import (
	"gotest.tools/assert"
	"os"
	"path"
	"testing"
)

func Dir(t *testing.T, testDir string, name string) string {
	t.Helper()

	p := path.Join(testDir, name)
	err := os.Mkdir(p, perm)
	assert.NilError(t, err)
	return p
}
