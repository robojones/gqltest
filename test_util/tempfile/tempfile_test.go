package tempfile

import (
	"github.com/robojones/gqltest/test_util/tempdir"
	"gotest.tools/assert"
	"io/ioutil"
	"os"
	"testing"
)

const (
	filename = "file.tmp"
	body     = "wazzup"
)

func TestCreate(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	p := Create(t, dir, filename, body)

	result, err := ioutil.ReadFile(p)
	assert.NilError(t, err)
	assert.Equal(t, string(result), body)
}

func TestRemove(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	p := Create(t, dir, filename, body)

	Remove(t, p)

	_, err := ioutil.ReadFile(p)
	assert.Assert(t, os.IsNotExist(err))
}
