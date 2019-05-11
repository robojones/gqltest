package reader

import (
	"github.com/vektah/gqlparser/ast"
)

type AllSources = map[string]*ast.Source

// read all tests.
func (r *Reader) Read(testdir string) map[string]*ast.Source {
	var (
		s   = make(AllSources)
		dir = r.ReadDir(testdir)
	)

	for _, name := range dir.Files {
		s[name] = r.ReadSource(name)
	}

	return s
}
