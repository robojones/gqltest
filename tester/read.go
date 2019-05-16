package tester

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/source"
	"github.com/robojones/gqltest/test"
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/parser"
)

func (t *Tester) read() ([]*test.Test, error) {
	files, err := source.ReadAll(t.config.TestGlob())
	if err != nil {
		panic(errors.Wrap(err, "read tests"))
	}

	var tests []*test.Test

	for _, src := range files {
		doc, err := parser.ParseQuery(src)
		if err != nil {
			return nil, errors.Wrapf(err, "Test file %s", src.Name)
		}

		ops := doc.Operations

		for _, op := range ops {
			query := &ast.QueryDocument{
				Position:   doc.Position,
				Operations: []*ast.OperationDefinition{op},
			}

			if err := t.validator.Validate(query); err != nil {
				return nil, err
			}

			t := test.NewTest(query)

			if err := t.ParseExpectations(); err != nil {
				return nil, err
			}

			tests = append(tests, t)
		}
	}

	return tests, nil
}
