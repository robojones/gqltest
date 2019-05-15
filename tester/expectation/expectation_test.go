package expectation

import (
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/parser"
	"gotest.tools/assert"
	"strings"
	"testing"
)

func fieldDirective(t *testing.T, directive string, params map[string]string) *ast.Directive {
	t.Helper()
	var args []string

	for key, val := range params {
		args = append(args, key+":"+val)
	}

	query := "query @" + directive + "(" + strings.Join(args, ",") + ") { asdf }"

	testSource := &ast.Source{
		Name:  "testfile.graphql",
		Input: query,
	}

	doc, err := parser.ParseQuery(testSource)

	if err != nil {
		panic(errors.WithStack(err))
	}

	op := doc.Operations[0]
	return op.Directives[0]
}

func TestFromDirective_MatchValueDirective(t *testing.T) {
	directive := fieldDirective(t, "expectString", map[string]string{
		"equal": "wazzup",
	})

	exp, err := FromDirective([]string{"asdf"}, directive)

	assert.NilError(t, err)
	e := exp.(*ValueExpectation)
	assert.Equal(t, e.value.(string), "wazzup")
}

func TestFromDirective_NoMatch(t *testing.T) {
	// expect doesn't match any directive
	directive := fieldDirective(t, "expect", map[string]string{
		"equal": "wazzup",
	})

	exp, err := FromDirective([]string{"asdf"}, directive)

	assert.NilError(t, err)
	assert.Assert(t, exp == nil)
}
