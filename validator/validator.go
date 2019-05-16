package validator

import (
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/validator"
)

type ValidatorConfig interface {
	SchemaGlob() string
	DirectivesFile() string
}

type ValidatorEnv interface {
	Branch() string
}

type Validator struct {
	conf   ValidatorConfig
	env    ValidatorEnv
	schema *ast.Schema
}

func NewValidator(c ValidatorConfig, env ValidatorEnv) (*Validator, error) {
	v := &Validator{
		conf: c,
		env:  env,
	}

	if err := v.readSchema(); err != nil {
		return nil, err
	}

	return v, nil
}

func (v *Validator) Validate(query *ast.QueryDocument) error {
	err := validator.Validate(v.schema, query)
	if err != nil {
		return err
	}
	return nil
}
