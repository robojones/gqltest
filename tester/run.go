package tester

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/poll"
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

	if err := poll.Poll(t.config.Endpoint(), t.config.StartTimeout()); err != nil {
		return err
	}

	for _, test := range tests {
		log.Printf("sending test body: %s", test.Body())
		v := request.NewVariables()
		p := request.NewPayload("Test", test.Body(), v)

		res, err := request.Send(t.config.Endpoint(), p, t.config.TestTimeout())

		if err != nil {
			return errors.Wrap(err, test.Operation().Position.Src.Name)
		}

		test.Verify(res)

		if len(test.Errors) != 0 {
			for _, err := range test.Errors {
				log.Println(err.Error())
			}
		}
	}

	log.Printf("successfully ran %d operations against %s.", len(tests), t.config.Endpoint())

	return nil
}
