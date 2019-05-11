package reader

import (
	"github.com/bmatcuk/doublestar"
	"github.com/vektah/gqlparser/ast"
	"os"
)

type AllSources = map[string]*ast.Source

// Read all files matching the glob.
func (r *Reader) Read(glob string) map[string]*ast.Source {
	var (
		s   = make(AllSources)
		files, err = doublestar.Glob(glob)
	)

	if err != nil {
		panic(err)
	}

	for _, name := range files {
		info, err := os.Stat(name)

		if err != nil {
			panic(err)
		}

		if !info.IsDir() {
			s[name] = r.ReadSource(name)
		}
	}

	return s
}
