package test

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/expectation"
	"github.com/vektah/gqlparser/ast"
)

func (e *Test) ParseExpectations() error {
	rootSelections := e.Operation.SelectionSet

	err := eachField(rootSelections, func(path []string, field *ast.Field) error {
		for _, directive := range field.Directives {
			exp, err := expectation.FromDirective(path, directive)

			if err != nil {
				return err
			}

			e.Expectations = append(e.Expectations, exp)
			e.Directives = append(e.Directives, directive.Position)
		}

		return nil
	}, []string{})

	return err
}

// eachSelection iterates through all selections recursively
func eachField(set ast.SelectionSet, fn func(path []string, field *ast.Field) error, path []string) error {
	for _, selection := range set {
		field, isField := selection.(*ast.Field)
		if !isField {
			panic(errors.Errorf("Selection %#v in %#v is not a field", selection, path))
		}

		path = append(path, field.Alias)

		if err := fn(path, field); err != nil {
			return err
		}

		if err := eachField(field.SelectionSet, fn, path); err != nil {
			return err
		}

		path = path[:len(path)-1]
	}

	return nil
}
