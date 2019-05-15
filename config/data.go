package config

const (
	DefaultSchemaGlob         = "**.graphqls"
	DefaultTestGlob           = "tests/**.graphql"
	DefaultStartTimeout int64 = 1000 * 10 // milliseconds
	DefaultTestTimeout  int64 = 1000 * 5  // milliseconds
)

// configData represents the contents of the gqltest.yml file.
type configData struct {
	// Endpoint specifies the URL to run the tests against.
	Endpoint string `yaml:"endpoint"`
	// SchameGlob matches the schema file names.
	SchemaGlob string `yaml:"schema"`
	// TestGlob matches the test file names.
	TestGlob string `yaml:"tests"`
	// DirectiveFile specifies the path to the file which contains the directives.
	// Contents be updated automatically when gqltest is run.
	DirectivesFile string `yaml:"directives"`
	// Initial timeout until the endpoint has to be online (in milliseconds).
	StartTimeout int64 `yaml:"startTimeout"`
	// Timeout for the tests.
	TestTimeout int64 `yaml:"testTimeout"`
}

func newDefaultConfigData() *configData {
	return &configData{
		SchemaGlob:   DefaultSchemaGlob,
		TestGlob:     DefaultTestGlob,
		StartTimeout: DefaultStartTimeout,
		TestTimeout:  DefaultTestTimeout,
	}
}
