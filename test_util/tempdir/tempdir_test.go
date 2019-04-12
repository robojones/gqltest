package tempdir

import (
	"gotest.tools/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestCreate(t *testing.T) {
	dir := Create(t)
	assert.Assert(t, dir != "")

	defer func () {
		assert.NilError(t, os.Remove(dir))
	}()

	_, err := ioutil.ReadDir(dir)
	assert.NilError(t, err)
}

func TestRemove(t *testing.T) {
	dir := Create(t)

	const filename = "file.tmp"

	p := path.Join(dir, filename)
	assert.NilError(t, ioutil.WriteFile(p, []byte{}, os.ModeTemporary))

	info, err := ioutil.ReadDir(dir)
	assert.NilError(t, err)
	assert.Equal(t, info[0].Name(), filename)

	Remove(t, dir)

	_, err = ioutil.ReadDir(dir)
	assert.Assert(t, os.IsNotExist(err))
}
