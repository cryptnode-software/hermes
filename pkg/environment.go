package pkg

import (
	commons "github.com/cryptnode-software/commons/pkg"
)

func NewEnv(logger commons.Logger) *Environment {
	return &Environment{
		logger,
	}
}

type Environment struct {
	Log commons.Logger
}
