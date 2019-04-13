package config

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"path"
)

func ReadConfig(wd WD) (*Config, error) {

	p := path.Join(string(wd), ConfigFileName)
	b, err := ioutil.ReadFile(p)

	if err != nil {
		return nil, errors.Wrap(err, "read config")
	}

	return parseConfig(b)
}
