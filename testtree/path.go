package testtree

import (
	"path"
)

type Path = []string

// NewPath creates a new path array from string path.
// p must not end with a /
func NewPath(p string) Path {
	dir, file := path.Split(path.Clean(p))

	if file == "" || file == "." {
		return Path{}
	}

	return append(NewPath(dir), file)
}
