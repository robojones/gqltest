package config

const ConfigFileName = "gqltest.yml"
const DefaultTestRoot = "tests"

type WD string

// Config represents the contents of the gqltest.yml file.
type Config struct {
	// Endpoint specifies the URL to run the tests against.
	Endpoint string
	// TestRoot is the root directory of your tests.
	TestRoot string
}

func newDefaultConfig() *Config {
	return &Config{
		TestRoot: DefaultTestRoot,
	}
}
