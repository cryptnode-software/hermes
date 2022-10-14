package pkg

import (
	"github.com/cryptnode-software/hermes/gorm"
	"github.com/google/uuid"
	api "go.buf.build/grpc/go/thenewlebowski/hermes/v1"
)

type event struct {
	*api.Event
}

func (event event) validate() (*gorm.Event, error) {
	if event.Text == "" {
		return nil, ErrEmptyMessageText
	}

	result := new(gorm.Event)

	if uuid, err := uuid.Parse(event.Id); err == nil {
		result.ID = uuid
	}

	result.Metadata = metadata{event.Metadata}.convert()
	result.Text = event.Text

	return result, nil
}

type metadata struct {
	value map[string]string
}

func (metadata metadata) convert() *gorm.Metadata {
	if len(metadata.value) == 0 {
		return nil
	}
	result := gorm.Metadata(metadata.value)
	return &result
}
