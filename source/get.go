package source

import (
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/ast"
	"io/ioutil"
	"net/http"
)

func Get(url string) (*ast.Source, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if resp.StatusCode != 200 {
		return nil, errors.WithStack(&StatusError{
			status: resp.Status,
			url:    url,
		})
	}

	b, err := ioutil.ReadAll(resp.Body)

	return &ast.Source{
		Name:  url,
		Input: string(b),
	}, err
}
