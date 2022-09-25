package pkg

import (
	"context"

	"github.com/cryptnode-software/hermes/model"
)

type Services struct {
	MessageService MessageService
	UserService    UserService
}

type UserService interface {
}

type MessageService interface {
	Delete(context.Context, *model.Message) error
	Save(context.Context, *model.Message) (*model.Message, error)
}
