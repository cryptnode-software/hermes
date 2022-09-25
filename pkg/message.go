package pkg

import (
	"github.com/cryptnode-software/hermes/model"
)

type MessageService interface {
	Delete(model.Message) (model.Message, error)
	Save(model.Message) (model.Message, error)
}
