package tester

import (
	"github.com/robojones/gqltest/config"
	"github.com/robojones/gqltest/source/reader"
	"github.com/robojones/gqltest/source/validator"
)

type Tester struct {
	config *config.Config
	reader *reader.Reader
	validator *validator.Validator
}

func NewTester(c *config.Config, r *reader.Reader, v *validator.Validator) *Tester {
	return &Tester{
		config: c,
		reader: r,
		validator: v,
	}
}
