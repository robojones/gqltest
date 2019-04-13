package config

const ConfigFileName = "gqltest.yml"
const DefaultTestRoot = "tests"

type WD string

// Config represents the contents of the gqltest.yml file.
type Config struct {
	// Endpoint specifies the URL to run the tests against.
	Endpoint string `yaml:"endpoint"`
	// TestRoot is the root directory of your tests.
	TestRoot string `yaml:"testRoot"`
}

func newDefaultConfig() *Config {
	return &Config{
		TestRoot: DefaultTestRoot,
	}
}
