package json

import "github.com/pkg/errors"

func DetectType(v *interface{}) string {
	switch (*v).(type) {
	case string:
		return "String"
	case float64:
		return "Number"
	case map[string]interface{}:
		return "Object"
	case []interface{}:
		return "Array"
	case bool:
		return "Boolean"
	case nil:
		return "null"
	default:
		panic(errors.Errorf("Unknown type %T of value %#v", v, v))
	}
}
