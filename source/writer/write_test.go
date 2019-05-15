package writer

import (
	"github.com/robojones/gqltest/source/reader"
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

	Write(s)

	r := reader.NewReader()
	actual := r.ReadSource(sourceName)

	assert.Equal(t, actual.Input, s.Input)
}
