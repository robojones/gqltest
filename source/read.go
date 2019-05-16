package source

import (
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/ast"
	"io/ioutil"
)

// Read a file and return its contents as a source struct.
func Read(filename string) (*ast.Source, error) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &ast.Source{
		Name:    filename,
		Input:   string(content),
		BuiltIn: false,
	}, nil
}
