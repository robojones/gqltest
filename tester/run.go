package tester

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/request"
)

func (t *Tester) Run() error {
	tests, err := t.Read()

	if err != nil {
		return err
	}

	for _, test := range tests {

		v := request.NewVariables()
		p := request.NewPayload("Test", string(test.Content()), v)
		r := request.NewRequest(t.config, p)
		_, err := request.Send(r)

		if err != nil {
			return errors.Wrap(err, test.Source.Name)
		}
	}

	return nil
}
