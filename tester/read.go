package tester

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/test"
	"github.com/vektah/gqlparser/parser"
)

func (t *Tester) Read() ([]*test.Test, error) {
	files := t.reader.Read(t.config.TestRoot())

	var tests []*test.Test

	for _, source := range files {
		doc, err := parser.ParseQuery(source)
		if err != nil {
			return nil, errors.Wrapf(err, "Test file %s", source.Name)
		}

		runes := []rune(source.Input)
		ops := doc.Operations

		for i, op := range ops {
			isLast := i == len(ops)-1
			var body string
			if isLast {
				body = string(runes[op.Position.Start:])
			} else {
				nextOp := ops[i+1]
				body = string(runes[op.Position.Start:nextOp.Position.Start])
			}
			tests = append(tests, test.NewTest(op, body))
		}
	}

	return tests, nil
}
