package config

import "path"

const ConfigFileName = "gqltest.yml"
const DefaultTestRoot = "tests"

type WD string

// Config
type Config struct {
	wd   string
	data *configData
}

func NewConfig(wd WD) (*Config, error) {
	d, err := readConfigData(wd)

	if err != nil {
		return nil, err
	}

	return &Config{
		wd: string(wd),
		data: d,
	}, nil
}

func (c *Config) Endpoint() string {
	return c.data.Endpoint
}

func (c *Config) TestRoot()string {
	return path.Join(c.wd, c.data.TestRoot)
}
