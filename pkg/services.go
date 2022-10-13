package pkg

import (
	"context"

	"github.com/cryptnode-software/hermes/model"
)

type UserService interface {
}

type Event interface {
	Delete(context.Context, *model.Event) error
	Save(context.Context, *model.Event) (*model.Event, error)
}
