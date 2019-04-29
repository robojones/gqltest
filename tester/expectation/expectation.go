package expectation

import (
	"github.com/robojones/gqltest/tester/request"
)

type Expectation = func(result *request.Result) error
