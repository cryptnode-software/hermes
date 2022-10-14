package gorm

import (
	commons "github.com/cryptnode-software/commons/pkg"
)

type Event struct {
	Metadata *Metadata
	commons.Model
	Text   string
	Author *User
}
