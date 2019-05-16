package test

import "github.com/robojones/gqltest/graphql"

func (t *Test) Body() string {
	return graphql.SerializeDocument(t.Document)
}
