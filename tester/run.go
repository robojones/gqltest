package tester

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/request"
	"log"
)

func (t *Tester) Run() error {
	tests, err := t.Read()

	if err != nil {
		return err
	}

	if len(tests) == 0 {
		return errors.New("No tests found.")
	}

	for _, test := range tests {
		log.Printf("sending test body: %s", test.Body())
		v := request.NewVariables()
		p := request.NewPayload("Test", test.Body(), v)
		r := request.NewRequest(t.config, p)
		_, err := request.Send(r)

		if err != nil {
			return errors.Wrap(err, test.Operation().Position.Src.Name)
		}

		test.Verify(r)
	}

	log.Printf("successfully ran %d operations against %s.", len(tests), t.config.Endpoint())

	return nil
}
