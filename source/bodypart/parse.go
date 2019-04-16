package bodypart

import (
	"github.com/graphql-go/graphql/language/parser"
	"github.com/graphql-go/graphql/language/source"
)

func ParseSource(s *source.Source) ([]*Part, error) {
	doc, err := parser.Parse(parser.ParseParams{
		Source:  s,
		Options: parser.ParseOptions{},
	})

	if err != nil {
		return nil, err
	}

	parts := make([]*Part, 0, len(doc.Definitions))

	for _, def := range doc.Definitions {
		parts = append(parts, NewPart(def, s))
	}

	return parts, nil
}
