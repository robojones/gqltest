package config

import (
	"github.com/pkg/errors"
)

func verifyConfig(c *configData) error {
	if c.Endpoint == "" {
		return errors.Wrap(VerificationError, "config must specify an endpoint")
	}

	return nil
}
