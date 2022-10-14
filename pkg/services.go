package pkg

import (
	"context"

	"github.com/cryptnode-software/hermes/gorm"
	"github.com/cryptnode-software/hermes/services/socket"
	api "go.buf.build/grpc/go/thenewlebowski/hermes/v1"
)

var services *Services = nil

type Services struct {
	event  EventService
	socket SocketService
	user   UserService
}

type UserService interface {
}

type EventService interface {
	Delete(context.Context, *gorm.Event) error
	Save(context.Context, *gorm.Event) (*gorm.Event, error)
}

type SocketService interface {
	Connect(ctx context.Context, resource string, socket socket.Socket) error
	Dispatch(ctx context.Context, resource string, events []*gorm.Event) (err error)
}

type Socket struct {
	stream api.Hermes_SubscribeServer
}

func (socket *Socket) Send(events []*gorm.Event) (err error) {
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
