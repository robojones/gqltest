package validator

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/source"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

func (v *Validator) readSchema() error {
	f, err := source.ReadAll(v.conf.SchemaGlob())
	if err != nil {
		panic(errors.Wrap(err, "read schema"))
	}

	files := make([]*ast.Source, len(f)+1)[:0]

	for _, file := range f {
		files = append(files, file)
	}

	d, err := v.directives()
	if err != nil {
		panic(err)
	}

	if v.conf.DirectivesFile() != "" {
		d.Name = v.conf.DirectivesFile()
		if source.Write(d) != nil {
			panic(errors.Wrap(err, "update test directives"))
		}
	}

	if _, containsDirectives := f[v.conf.DirectivesFile()]; !containsDirectives {
		files = append(files, d)
	}

	schema, gqlErr := gqlparser.LoadSchema(files...)
	if gqlErr != nil {
		return errors.WithStack(gqlErr)
	}

	v.schema = schema

	return nil
}
