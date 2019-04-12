package tester

import "github.com/robojones/gqltest/source/reader"

type Tester struct {
	reader *reader.Reader
}

func NewTester(r *reader.Reader) *Tester {
	return &Tester{
		reader: r,
	}
}
