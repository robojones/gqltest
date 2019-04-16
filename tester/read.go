package tester

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/source/bodypart"
)

func (t *Tester) Read() {
	s := t.reader.Read(t.config.TestRoot())

	for _, source := range s {
		parts :=

		if err != nil {
			return errors.Wrapf(err, "Test file %s", source.Name)
		}
	}

	return nil
}
