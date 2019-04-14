package reader

import (
	"github.com/graphql-go/graphql/language/source"
)

const TestFile = "test.graphql"

// Read all tests.
func (r *Reader) Read(testdir string) []*source.Source {
	var (
		s   []*source.Source
		dir = r.ReadDir(testdir)
	)

	for _, name := range dir.Files {
		s = append(s, r.ReadSource(name))
	}

	return s
}
