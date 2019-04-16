package tester

import (
	"github.com/robojones/gqltest/source/bodypart"
)

func (t *Tester) Read() ([]*bodypart.Part, error) {

	var p []*bodypart.Part
	s := t.reader.Read(t.config.TestRoot())

	for _, source := range s {
		parts, err := bodypart.ParseSource(source)

		if err != nil {
			//return errors.Wrapf(err, "Test file %s", source.Name)
		}

		p = append(p, parts...)
	}

	return p, nil
}
