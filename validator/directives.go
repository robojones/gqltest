package validator

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/source"
	"github.com/vektah/gqlparser/ast"
)

const (
	Filename    = "directives/directives.graphqls"
	URLTemplate = "https://raw.githubusercontent.com/robojones/gqltest/%s/" + Filename
)

func (v *Validator) directives() (*ast.Source, error) {
	branch := v.env.Branch()
	url := fmt.Sprintf(URLTemplate, branch)

	directives, err := source.Get(url)

	if err != nil {
		return nil, errors.Wrap(err, "get directives")
	}

	return directives, nil
}
