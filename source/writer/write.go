package writer

import (
	"github.com/vektah/gqlparser/ast"
	"io/ioutil"
)

const WritePerm = 0644

func Write(s *ast.Source) {
	err := ioutil.WriteFile(s.Name, []byte(s.Input), WritePerm)

	if err != nil {
		panic(err)
	}
}
