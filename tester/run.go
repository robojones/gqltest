package tester

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/request"
)

func (t *Tester) Run() error {
	s := t.reader.Read(t.config.TestRoot())

	for _, source := range s {
		v := request.NewVariables()
		p := request.NewPayload("Test", string(source.Body), v)
		r := request.NewRequest(t.config, p)
		_, err := request.Send(r)

		if err != nil {
			return errors.Wrapf(err, "Test file %s", source.Name)
		}
	}

	return nil
}
