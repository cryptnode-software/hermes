package pkg

import (
	"github.com/cryptnode-software/hermes/model"
	"github.com/golang/protobuf/ptypes"
	api "go.buf.build/grpc/go/thenewlebowski/hermes/v1"
)

type cmessage struct {
	message *model.Message
}

func (message cmessage) convert() (*api.Message, error) {
	if message.message == nil {
		return nil, nil
	}

	created, err := ptypes.TimestampProto(message.message.CreatedAt)
	if err != nil {
		return nil, err
	}

	updated, err := ptypes.TimestampProto(message.message.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &api.Message{
		Metadata: cmetadata{message.message.Metadata}.convert(),
		Id:       message.message.ID.String(),
		Text:     message.message.Text,
		Created:  created,
		Updated:  updated,
	}, nil
}

type cmetadata struct {
	metadata *model.Metadata
}

func (metadata cmetadata) convert() map[string]string {
	if metadata.metadata == nil {
		return nil
	}

	return *metadata.metadata
}
