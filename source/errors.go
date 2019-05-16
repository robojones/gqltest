package source

import "fmt"

type StatusError struct {
	status string
	url    string
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("status %s GET %s", e.status, e.url)
}
