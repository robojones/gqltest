package config

import "log"

func verifyConfig(c *Config) error {
	if c.Endpoint == "" {
		log.Fatalf("config must specify an endpoint")
	}

	return nil
}
