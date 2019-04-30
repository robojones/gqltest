package expectation

import (
	"github.com/robojones/gqltest/tester/request"
)

type Expectation interface {
	Check(result *request.Result) error
}
