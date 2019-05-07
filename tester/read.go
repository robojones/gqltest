package tester

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/test"
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/parser"
)

func (t *Tester) Read() ([]*test.Test, error) {
	files := t.reader.Read(t.config.TestRoot())

	var tests []*test.Test

	for _, source := range files {
		doc, err := parser.ParseQuery(source)
		if err != nil {
			return nil, errors.Wrapf(err, "Test file %s", source.Name)
		}

		ops := doc.Operations

		for _, op := range ops {
			doc := &ast.QueryDocument{
				Position:   doc.Position,
				Operations: []*ast.OperationDefinition{op},
			}

			tests = append(tests, test.NewTest(doc))
		}
	}

	return tests, nil
}
