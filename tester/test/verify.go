package test

import (
	"github.com/robojones/gqltest/tester/request"
	"github.com/robojones/gqltest/tester/testerror"
)

func (t *Test) Verify(result request.Result) {
	for _, exp := range t.Expectations {
		err := exp.Check(result)

		if err != nil {
			testErr := testerror.NewTestError(err, t.Document.Operations[0], exp.Directive())
			t.Errors = append(t.Errors, testErr)
		}
	}
}
