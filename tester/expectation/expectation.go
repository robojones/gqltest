package expectation

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/tester/request"
	"github.com/vektah/gqlparser/ast"
)

type Expectation interface {
	Directive() *ast.Directive
	Check(result request.Result) error
}

func FromDirective(path []string, directive *ast.Directive) (Expectation, error) {
	switch directive.Name {
	case "expect":
		if len(directive.Arguments) < 1 {
			return nil, errors.New("missing argument v")
		}
		value := directive.Arguments[0].Value

		raw := value.Raw
		var v interface{}

		switch value.Kind {
		case ast.Variable:
			return nil, errors.New("variables not supported")
		case ast.FloatValue, ast.IntValue, ast.BooleanValue, ast.ListValue, ast.ObjectValue, ast.NullValue:
			err := json.Unmarshal([]byte(raw), &v)
			if err != nil {
				panic(errors.Wrap(err, "parsing argument value"))
			}
			break
		case ast.StringValue, ast.EnumValue:
			v = raw
			break
		}

		dataPath := append([]string{"data"}, path...)
		return NewValueExpectation(directive, dataPath, v), nil
	}

	return nil, nil
}
