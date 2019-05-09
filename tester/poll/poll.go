package poll

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/request"
	"net"
	"net/url"
	"time"
)

const dial = "dial"

func IsDialError(err error) bool {
	cause := errors.Cause(err)
	urlErr, ok := cause.(*url.Error)
	if !ok {
		return false
	}
	opErr, ok := urlErr.Err.(*net.OpError)
	if !ok {
		return false
	}
	return opErr.Op == dial
}

func Poll(endpoint string, timeout time.Duration) error {
	query := "query {__schema {queryType {name}}}"

	p := request.NewPayload("pollUntilOnline", query, request.NewVariables())

	for {
		_, err := request.Send(endpoint, p, timeout)

		if err == nil {
			return nil
		} else if !IsDialError(err) {
			return errors.Wrap(err, "initial connect")
		}
	}
}
