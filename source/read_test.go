package source

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/test_util/tempdir"
	"gotest.tools/assert"
	"os"
	"testing"
)

const name = "testfile.json"

func TestReader_ReadSource(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	const body = "{}\n"
	p := tempdir.File(t, dir, name, body)

	source, err := Read(p)

	assert.NilError(t, err)
	assert.Equal(t, source.Name, p)
	assert.DeepEqual(t, source.Input, body)
}

func TestReader_ReadSource_Panic(t *testing.T) {
	_, err := Read("fileDoesNotExist.json")
	assert.Assert(t, os.IsNotExist(errors.Cause(err)))
}
