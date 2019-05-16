package env

import (
	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

const DefaultBranch = "master"

type Env struct {
	data *data
}

func NewEnv() *Env {
	e := &data{}
	if err := env.Parse(e); err != nil {
		panic(errors.WithStack(err))
	}

	return &Env{
		data: e,
	}
}

func (e *Env) Branch() string {
	if e.data.Branch != "" {
		return e.data.Branch
	}
	return DefaultBranch
}
