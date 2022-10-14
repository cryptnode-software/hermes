package pkg

import "errors"

var (
	ErrEmptyMessageText = errors.New("message received with empty text")
)
