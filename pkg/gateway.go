package pkg

import (
	"context"

	commons "github.com/cryptnode-software/commons/pkg"
	"github.com/cryptnode-software/hermes/gorm"
	api "go.buf.build/grpc/go/thenewlebowski/hermes/v1"
)

func NewGateway() (result *Gateway, err error) {
	result = new(Gateway)
	if result.log, err = commons.GetLogger(); err != nil {
		return nil, err
	}
	return
}

type Gateway struct {
	log    commons.Logger
	event  EventService
	socket SocketService
	user   UserService
}

func (gateway *Gateway) SetSocket(socket SocketService) error {
	gateway.socket = socket
	return nil
}

func (gateway *Gateway) SetEvent(event EventService) error {
	gateway.event = event
	return nil
}

func (gateway *Gateway) SetUser(user UserService) error {
	gateway.user = user
	return nil
}

func (gateway *Gateway) Subscribe(request *api.SubscribeRequest, server api.Hermes_SubscribeServer) (err error) {
	ctx := context.Background()

	if gateway.socket != nil {

		socket := &Socket{
			server,
		}

		if err = gateway.socket.Connect(ctx, request.Resource, socket); err != nil {
			gateway.log.Error("error during socket connection", "err", err)
			return
		}
	}

	return
}

func (gateway *Gateway) Dispatch(ctx context.Context, request *api.DispatchRequest) (response *api.DispatchResponse, err error) {
	event, err := event{request.Event}.validate()
	if err != nil {
		gateway.log.Error("save message error", "err", err)
		return
	}

	if gateway.socket != nil {
		if request.Resource != "" {
			gateway.socket.Dispatch(ctx, request.Resource, []*gorm.Event{event})
		}
	}

	if gateway.event != nil {
		event, err = gateway.event.Save(ctx, event)
		if err != nil {
			gateway.log.Error("save message error", "err", err)
			return
		}
	}

	response = new(api.DispatchResponse)
	response.Event, err = revent{event}.convert()
	if err != nil {
		gateway.log.Error("save message error", "err", err)
		return
	}

	return
}

func (gateway *Gateway) DeleteEvent(ctx context.Context, request *api.Event) (response *api.Event, err error) {
	event, err := event{request}.validate()
	if err != nil {
		gateway.log.Error("save message error", "err", err)
		return
	}

	if gateway.event != nil {
		err = gateway.event.Delete(ctx, event)
		if err != nil {
			gateway.log.Error("delete message error", "err", err)
			return
		}
	}

	return
}
