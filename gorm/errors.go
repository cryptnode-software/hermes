package gorm

import "errors"

var (
	ErrNoDatabaseSet = errors.New("no database is set prior to retrieval")
)
