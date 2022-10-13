package pkg

import (
	"context"

	"github.com/cryptnode-software/hermes/model"
	api "go.buf.build/grpc/go/thenewlebowski/hermes/v1"
)

type SocketService interface {
	Connect(resource string, socket Socket) error
	Dispatch(ctx context.Context, resource string, events []*model.Event) (err error)
}

type Socket interface {
	Send([]*model.Event) error
}

type socket struct {
	stream api.Hermes_SubscribeServer
}

func (socket *socket) Send(events []*model.Event) (err error) {
	result := make([]*api.Event, len(events))

	for i := range events {
		result[i], err = revent{events[i]}.convert()
		if err != nil {
			return err
		}
	}

	socket.stream.Send(&api.SubscribeResponse{
		Events: result,
	})

	return nil
}
