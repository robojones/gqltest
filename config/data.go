package config

const (
	DefaultTestRoot       = "tests"
	DefaultStartupTimeout = 1000 * 10 // milliseconds
	DefaultTestTimeout    = 1000 * 5  // milliseconds
)

// configData represents the contents of the gqltest.yml file.
type configData struct {
	// Endpoint specifies the URL to run the tests against.
	Endpoint string `yaml:"endpoint"`
	// TestRoot is the root directory of your tests.
	TestRoot string `yaml:"testRoot"`
	// Initial timeout until the endpoint has to be online (in milliseconds).
	StartTimeout int64 `yaml:"startTimeout"`
	// Timeout for the tests.
	TestTimeout int64 `yaml:"testTimeout"`
}

func newDefaultConfigData() *configData {
	return &configData{
		TestRoot:     DefaultTestRoot,
		StartTimeout: DefaultStartupTimeout,
		TestTimeout:  DefaultTestTimeout,
	}
}
