package validator

import (
	"github.com/robojones/gqltest/source/directives"
	"github.com/robojones/gqltest/source/reader"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/validator"
	"log"
)

type ValidatorConfig interface {
	SchemaGlob() string
}

type Validator struct {
	schema *ast.Schema
}

func NewValidator(c ValidatorConfig, r *reader.Reader) (*Validator, error) {
	f := r.Read(c.SchemaGlob())
	files := make([]*ast.Source, len(f)+1)[:0]

	for _, file := range f {
		files = append(files, file)
	}

	d, err := directives.Get()
	if err != nil {
		return nil, err
	}
	files = append(files, d)

	schema, gqlErr := gqlparser.LoadSchema(files...)
	if gqlErr != nil {
		log.Printf("LoadSchema returned error")
		return nil, gqlErr
	}

	return &Validator{
		schema: schema,
	}, nil
}

func (v *Validator) Validate(query *ast.QueryDocument) error {
	err := validator.Validate(v.schema, query)
	if err != nil {
		return err
	}
	return nil
}
