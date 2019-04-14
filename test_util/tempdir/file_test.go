package tempdir

import (
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func TestFile(t *testing.T) {
	const (
		filename = "file.tmp"
		body     = "wazzup"
	)

	dir := Create(t)
	defer Remove(t, dir)

	p := File(t, dir, filename, body)

	result, err := ioutil.ReadFile(p)
	assert.NilError(t, err)
	assert.Equal(t, string(result), body)
}
