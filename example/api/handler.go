package api

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/example/api/query"
	"io/ioutil"
)

func NewHandler(query *query.Resolver) *relay.Handler {
	s, err := ioutil.ReadFile("example/schema.graphqls")

	if err != nil {
		panic(errors.Wrap(err, "Example api"))
	}

	schema := graphql.MustParseSchema(string(s), query)

	return &relay.Handler{Schema: schema}
}
