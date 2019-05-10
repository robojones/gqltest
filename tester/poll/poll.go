package poll

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/request"
	"net"
	"net/url"
	"time"
)

const (
	dial  = "dial"
	op    = "pollUntilOnline"
	query = "query {__schema {queryType {name}}}"
)

const interval = 100

func Poll(endpoint string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)

	p := request.NewPayload(op, query, request.NewVariables())

	for {
		t := deadline.Sub(time.Now())

		if t <= 0 {
			// because 0 means no timeout in the http client
			t = 1
		}

		_, err := request.Send(endpoint, p, t)

		if err == nil {
			return nil
		} else if !IsDialError(err) || IsTimeoutError(err) {
			return errors.Wrap(err, "initial connect")
		}
		time.Sleep(interval * time.Millisecond)
	}
}

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

func IsTimeoutError(err error) bool {
	cause := errors.Cause(err)
	netErr, ok := cause.(net.Error)
	if !ok {
		return false
	}
	return netErr.Timeout()
}
