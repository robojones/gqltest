package tester

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/source"
	"github.com/robojones/gqltest/tester/request"
	"log"
)

func (t *Tester) Run() error {
	tests, err := t.Read()

	if err != nil {
		return err
	}

	for _, test := range tests {
		v := request.NewVariables()
		p := request.NewPayload("Test", source.Content(test.Position), v)
		r := request.NewRequest(t.config, p)
		_, err := request.Send(r)

		if err != nil {
			return errors.Wrapf(err, "%s: %s", test.Position.Src.Name, test.Name)
		}
	}

	log.Printf("successfully ran %d operations against %s.", len(tests), t.config.Endpoint())

	return nil
}
