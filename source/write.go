package source

import (
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/ast"
	"io/ioutil"
)

const WritePerm = 0644

// Write a source to the disk.
func Write(s *ast.Source) error {
	err := ioutil.WriteFile(s.Name, []byte(s.Input), WritePerm)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
