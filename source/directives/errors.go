package directives

import "fmt"

type StatusError struct {
	statusCode int
	status     string
	url        string
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("status %d (%s) GET %s", e.statusCode, e.status, e.url)
}
