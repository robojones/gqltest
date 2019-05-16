package tester

import (
	"github.com/robojones/gqltest/validator"
	"time"
)

type Tester struct {
	config    TesterConfig
	validator *validator.Validator
}

type TesterConfig interface {
	Endpoint() string
	TestGlob() string
	StartTimeout() time.Duration
	TestTimeout() time.Duration
}

func NewTester(c TesterConfig, v *validator.Validator) *Tester {
	return &Tester{
		config:    c,
		validator: v,
	}
}
