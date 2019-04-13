package tester

import (
	"github.com/robojones/gqltest/tester/request"
)

func (t *Tester) Run() error {
	s := t.reader.Read(t.config.TestRoot())

	v := request.NewVariables()
	p := request.NewPayload("Test", string(s.Body), v)
	r := request.NewRequest(t.config, p)

	_, err := request.Send(r)

	return err
}
