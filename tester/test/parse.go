package test

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/expectation"
	"github.com/vektah/gqlparser/ast"
)

func (t *Test) ParseExpectations() error {
	rootSelections := t.Document.Operations[0].SelectionSet

	err := t.parseEachField(rootSelections, []string{})

	return err
}

func (t *Test) parseField(path []string, field *ast.Field) error {
	var unknown ast.DirectiveList

	for _, directive := range field.Directives {
		exp, err := expectation.FromDirective(path, directive)

		if err != nil {
			return err
		}

		if exp == nil {
			unknown = append(unknown, directive)
		} else {
			t.Expectations = append(t.Expectations, exp)
		}

	}

	field.Directives = unknown

	return nil
}

func (t *Test) parseEachField(set ast.SelectionSet, path []string) error {
	for _, selection := range set {
		field, isField := selection.(*ast.Field)
		if !isField {
			panic(errors.Errorf("Selection %#v in %#v is not a field", selection, path))
		}

		path = append(path, field.Alias)

		if err := t.parseField(path, field); err != nil {
			return err
		}

		if err := t.parseEachField(field.SelectionSet, path); err != nil {
			return err
		}

		path = path[:len(path)-1]
	}

	return nil
}
