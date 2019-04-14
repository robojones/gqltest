package tempdir

import (
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func TestDir(t *testing.T) {
	const (
		subDirName = "file"
		body       = "wazzup"
	)

	dir := Create(t)
	defer Remove(t, dir)

	p := Dir(t, dir, subDirName)

	_, err := ioutil.ReadDir(p)
	assert.NilError(t, err)
}
