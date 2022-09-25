package pkg

import (
	"github.com/cryptnode-software/hermes/model"
	"github.com/google/uuid"
	api "go.buf.build/grpc/go/thenewlebowski/hermes/v1"
)

type message struct {
	*api.Message
}

func (message message) validate() (*model.Message, error) {
	if message.Text == "" {
		return nil, ErrEmptyMessageText
	}

	result := new(model.Message)

	if uuid, err := uuid.Parse(message.Id); err == nil {
		result.ID = uuid
	}

	result.Metadata = metadata{message.Metadata}.convert()
	result.Text = message.Text

	return result, nil
}

type metadata struct {
	value map[string]string
}

func (metadata metadata) convert() *model.Metadata {
	if len(metadata.value) == 0 {
		return nil
	}
	result := model.Metadata(metadata.value)
	return &result
}
