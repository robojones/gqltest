package reader

import (
	"github.com/graphql-go/graphql/language/source"
	"path"
)

const TestFile = "test.graphql"

// Read all tests.
func (r *Reader) Read(testdir string) *source.Source {
	p := path.Join(testdir, TestFile)
	return r.ReadSource(p)
}
