package model

import (
	commons "github.com/cryptnode-software/commons/pkg"
)

type Message struct {
	Metadata *Metadata
	commons.Model
	Text   string
	Author *User
}
