package config

import (
	"github.com/pkg/errors"
)

func verifyConfig(c *Config) error {
	if c.Endpoint == "" {
		return errors.Wrap(VerificationError, "config must specify an endpoint")
	}

	return nil
}
