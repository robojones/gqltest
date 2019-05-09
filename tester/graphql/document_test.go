package graphql

import (
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/parser"
	"gotest.tools/assert"
	"testing"
)

var testDocs = []string{`
query Cool { field }
`, `
mutation Such { wow }
`, `
query Wazzup($someVar: String!, $other: Int) @directive(arg: "stringArg") {
	alias: field
}`, `
query ($someVar: String) @directive {
	foo {
		nestedField @firstDirective(value: $someVar) @secondDirective(foo: "string")
	}
}`, `
query {
	foo {
		...on Bar {
			baz
		}
	}
}
`, `
fragment Cool on SomeType @withDirective {
	selection
}
`, `
query {
	foo {
		... {
			baz
		}
	}
}
`,
}

func TestSerializeDocument(t *testing.T) {
	for _, testInput := range testDocs {
		original, err := parser.ParseQuery(&ast.Source{
			Name:  "testQuery.graphql",
			Input: testInput,
		})

		if err != nil {
			panic(err)
		}

		generated := SerializeDocument(original)

		result, err := parser.ParseQuery(&ast.Source{
			Name:  "result.graphql",
			Input: generated,
		})

		if err != nil {
			panic(err)
		}

		assert.Equal(t, ast.Dump(original), ast.Dump(result))
	}
}
