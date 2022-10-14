package gorm

import (
	commons "github.com/cryptnode-software/commons/pkg"
)

type User struct {
	Metadata  *Metadata
	FirstName string
	LastName  string
	Username  string
	Email     string
	commons.Model
}
