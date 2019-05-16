package source

import (
	"github.com/bmatcuk/doublestar"
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/ast"
	"os"
)

type AllSources = map[string]*ast.Source

// ReadAll all files matching the glob.
func ReadAll(glob string) (map[string]*ast.Source, error) {
	s := make(AllSources)

	files, err := doublestar.Glob(glob)
	if err != nil {
		return nil, errors.Wrapf(err, "resolve \"%s\"", glob)
	}

	for _, name := range files {
		info, err := os.Stat(name)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		if !info.IsDir() {
			s[name], err = Read(name)
			if err != nil {
				return nil, errors.WithStack(err)
			}
		}
	}

	return s, nil
}
