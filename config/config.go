package config

import (
	"github.com/pkg/errors"
	"path"
)

const ConfigFileName = "gqltest.yml"
const DefaultTestRoot = "tests"

type WorkinDirectoryName string

// Config
type Config struct {
	wd   string
	data *configData
}

func NewConfig(wd WorkinDirectoryName) (*Config, error) {
	d, err := readConfigData(wd)

	if err != nil {
		return nil, errors.Wrap(err, string(wd))
	}

	return &Config{
		wd:   string(wd),
		data: d,
	}, nil
}

func (c *Config) Endpoint() string {
	return c.data.Endpoint
}

func (c *Config) TestRoot() string {
	return path.Join(c.wd, c.data.TestRoot)
}
