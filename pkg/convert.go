package pkg

import (
	"github.com/cryptnode-software/hermes/model"
	"github.com/golang/protobuf/ptypes"
	api "go.buf.build/grpc/go/thenewlebowski/hermes/v1"
)

type revent struct {
	event *model.Event
}

func (event revent) convert() (*api.Event, error) {
	if event.event == nil {
		return nil, nil
	}

	created, err := ptypes.TimestampProto(event.event.CreatedAt)
	if err != nil {
		return nil, err
	}

	updated, err := ptypes.TimestampProto(event.event.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &api.Event{
		Metadata: cmetadata{event.event.Metadata}.convert(),
		Id:       event.event.ID.String(),
		Text:     event.event.Text,
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
