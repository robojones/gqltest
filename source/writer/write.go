package writer

import (
	"github.com/vektah/gqlparser/ast"
	"io/ioutil"
)

const WritePerm = 0644

func Write(s *ast.Source) error {
	return ioutil.WriteFile(s.Name, []byte(s.Input), WritePerm)
}
