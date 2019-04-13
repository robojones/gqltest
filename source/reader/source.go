package reader

import (
	"github.com/graphql-go/graphql/language/source"
	"github.com/pkg/errors"
	"io/ioutil"
)

func (r *Reader) readSource(filename string) *source.Source {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(errors.Wrap(err, filename))
	}

	return &source.Source{
		Name: filename,
		Body: content,
	}
}
