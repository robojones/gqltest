package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func parseConfig(b []byte) (*configData, error) {
	c := newDefaultConfigData()
	err := yaml.UnmarshalStrict(b, c)

	if err != nil {
		return nil, errors.Wrap(err, "parse config")
	}

	return c, verifyConfig(c)
}
