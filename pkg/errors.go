package pkg

import "errors"

var (
	ErrNoDatabaseSet    = errors.New("no database is set prior to retrevial")
	ErrEmptyMessageText = errors.New("message recieved with empty text")
)
