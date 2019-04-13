package tester

import (
	"github.com/robojones/gqltest/config"
	"github.com/robojones/gqltest/source/reader"
)

type Tester struct {
	config *config.Config
	reader *reader.Reader
}

func NewTester(c *config.Config, r *reader.Reader) *Tester {
	return &Tester{
		config: c,
		reader: r,
	}
}
