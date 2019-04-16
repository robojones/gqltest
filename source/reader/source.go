package reader

import (
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/ast"
	"io/ioutil"
)

func (r *Reader) ReadSource(filename string) *ast.Source {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(errors.Wrap(err, filename))
	}

	return &ast.Source{
		Name:    filename,
		Input:   string(content),
		BuiltIn: false,
	}
}
