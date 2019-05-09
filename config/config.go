package config

import (
	"github.com/pkg/errors"
	"path"
	"time"
)

const (
	ConfigFileName        = "gqltest.yml"
)

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

func (c *Config) StartTimeout() time.Duration {
	return time.Duration(c.data.StartTimeout) * time.Millisecond
}

func (c *Config) TestTimeout() time.Duration {
	return time.Duration(c.data.TestTimeout) * time.Millisecond
}
