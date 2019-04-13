package reader

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/test_util/tempdir"
	"github.com/robojones/gqltest/test_util/tempfile"
	"gotest.tools/assert"
	"os"
	"testing"
)

const name = "testfile.json"

func TestReadSource(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	const body = "{}\n"
	p := tempfile.Create(t, dir, name, body)

	reader := NewReader()

	source := reader.readSource(p)

	assert.Equal(t, source.Name, p)
	assert.DeepEqual(t, string(source.Body), body)
}

func TestReadSourcePanic(t *testing.T) {
	defer func() {
		err := recover().(error)
		assert.Assert(t, os.IsNotExist(errors.Cause(err)))
	}()

	reader := NewReader()

	reader.readSource("fileDoesNotExist.json")
}
