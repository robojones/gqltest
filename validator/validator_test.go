package validator

import (
	"github.com/robojones/gqltest/env"
	"github.com/robojones/gqltest/source"
	"github.com/robojones/gqltest/test_util/tempdir"
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/parser"
	"gotest.tools/assert"
	"path"
	"testing"
)

const (
	graphqlsGlob     = "**/*.graphqls"
	testSchemaName   = "schema.graphqls"
	testQueryName    = "testQuery.graphql"
	invalidQueryName = "invalidQuery.graphql"
	testSchema       = `
type Query {
	user: User!
}

type User {
	id: ID!
	email: String!

	name: String!
	createdAt: String!
	updatedAt: String!
}
`
	testQuery = `
query {
	user {
		email @expectString(equal: "cool@email.com")
	}
}
`
	invalidQuery = `
query {
	user {
		# age does not exist on type user
		age
	}
}
`
)

type testConfig struct {
	schemaGlob string
	directives string
}

func (c *testConfig) SchemaGlob() string {
	return c.schemaGlob
}

func (c *testConfig) DirectivesFile() string {
	return c.directives
}

func TestNewValidator(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	tempdir.File(t, dir, testSchemaName, testSchema)

	c := &testConfig{
		schemaGlob: path.Join(dir, graphqlsGlob),
	}

	validator, err := NewValidator(c, env.NewEnv())
	assert.NilError(t, err)

	doc, gqlErr := parser.ParseQuery(&ast.Source{
		Name:  testQueryName,
		Input: testQuery,
	})
	assert.Assert(t, gqlErr == nil, gqlErr.Error())

	err = validator.Validate(doc)
	assert.NilError(t, err)
}

func TestNewValidatorWriteDirectives(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	directives := path.Join(dir, "directives.graphql")

	c := &testConfig{
		schemaGlob: path.Join(dir, graphqlsGlob),
		directives: directives,
	}

	_, err := NewValidator(c, env.NewEnv())
	assert.NilError(t, err)

	d, err := source.Read(directives)
	assert.NilError(t, err)

	assert.Equal(t, d.Name, directives)
	assert.Assert(t, d.Input != "")
}

func TestNewValidatorInvalidQuery(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	tempdir.File(t, dir, testSchemaName, testSchema)

	c := &testConfig{
		schemaGlob: path.Join(dir, graphqlsGlob),
	}

	validator, err := NewValidator(c, env.NewEnv())
	assert.NilError(t, err)

	doc, gqlErr := parser.ParseQuery(&ast.Source{
		Name:  invalidQueryName,
		Input: invalidQuery,
	})
	assert.Assert(t, gqlErr == nil, gqlErr.Error())

	err = validator.Validate(doc)
	assert.ErrorContains(t, err, "age")
}
