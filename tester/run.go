package tester

import "github.com/robojones/gqltest/tester/request"

const testdir = "tests"

func (t *Tester) Run() error {
	s := t.reader.Read(testdir)

	v := request.NewVariables()
	p := request.NewPayload("Test", string(s.Body), v)
	r := request.NewRequest(p)

	_, err := request.Send(r)

	return err
}
