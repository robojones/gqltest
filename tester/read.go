package tester

import (
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/parser"
)

func (t *Tester) Read() ([]*ast.OperationDefinition, error) {
	files := t.reader.Read(t.config.TestRoot())

	var ops []*ast.OperationDefinition

	for _, source := range files {
		doc, err := parser.ParseQuery(source)

		if err != nil {
			return nil, errors.Wrapf(err, "Test file %s", source.Name)
		}

		ops = append(ops, doc.Operations...)
	}

	return ops, nil
}
