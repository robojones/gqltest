package source

import (
	"github.com/robojones/gqltest/test_util/tempdir"
	"github.com/vektah/gqlparser/ast"
	"gotest.tools/assert"
	"path"
	"testing"
)

func TestWrite(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	const filename = "exampleFile.txt"
	const sourceContent = "this is some cool content"
	sourceName := path.Join(dir, filename)

	s := &ast.Source{
		Name:  sourceName,
		Input: sourceContent,
	}

	assert.NilError(t, Write(s))

	actual, err := Read(sourceName)
	assert.NilError(t, err)

	assert.Equal(t, actual.Input, s.Input)
}

func TestWriteOverwrite(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	const filename = "exampleFile.txt"

	s := &ast.Source{
		Name: path.Join(dir, filename),
	}

	assert.NilError(t, Write(s))
	assert.NilError(t, Write(s))
}
