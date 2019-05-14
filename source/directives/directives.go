package directives

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/ast"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	DefaultBranch = "master"
	BranchEnvKey  = "GQLTEST_BRANCH"
	Filename      = "directives/directives.graphqls"
	URLTemplate   = "https://raw.githubusercontent.com/robojones/gqltest/%s/" + Filename
)

func getBranch() string {
	branch, ok := os.LookupEnv(BranchEnvKey)

	if !ok {
		branch = DefaultBranch
	}

	return branch
}

func Get() (*ast.Source, error) {
	branch := getBranch()
	url := fmt.Sprintf(URLTemplate, branch)

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.Wrapf(err, "request directives for branch \"%s\"", branch)
	}

	if resp.StatusCode != 200 {
		return nil, errors.Wrapf(&StatusError{
			statusCode: resp.StatusCode,
			status:     resp.Status,
			url:        url,
		},
			"request directives for branch \"%s\"",
			branch,
		)
	}

	b, err := ioutil.ReadAll(resp.Body)

	return &ast.Source{
		Name:  URLTemplate,
		Input: string(b),
	}, err
}
