package config

// configData represents the contents of the gqltest.yml file.
type configData struct {
	// Endpoint specifies the URL to run the tests against.
	Endpoint string `yaml:"endpoint"`
	// TestRoot is the root directory of your tests.
	TestRoot string `yaml:"testRoot"`
}

func newDefaultConfigData() *configData {
	return &configData{
		TestRoot: DefaultTestRoot,
	}
}
