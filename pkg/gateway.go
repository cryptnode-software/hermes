package pkg

import (
	"context"

	commons "github.com/cryptnode-software/commons/pkg"
	"github.com/cryptnode-software/hermes/model"
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
	log         commons.Logger
	UserService UserService
	SocketService
	Event Event
}

func (gateway *Gateway) Subscribe(request *api.SubscribeRequest, server api.Hermes_SubscribeServer) error {
	socket := &socket{
		server,
	}
	return gateway.SocketService.Connect(request.Resource, socket)
}

func (gateway *Gateway) Dispatch(ctx context.Context, request *api.DispatchRequest) (response *api.DispatchResponse, err error) {
	event, err := event{request.Event}.validate()
	if err != nil {
		gateway.log.Error("save message error", "err", err)
		return
	}

	if request.Resource != "" {
		gateway.SocketService.Dispatch(ctx, request.Resource, []*model.Event{event})
	}

	event, err = gateway.Event.Save(ctx, event)
	if err != nil {
		gateway.log.Error("save message error", "err", err)
		return
	}
	response = new(api.DispatchResponse)
	response.Event, err = revent{event}.convert()
	if err != nil {
		gateway.log.Error("save message error", "err", err)
		return
	}

	return
}

func (gateway *Gateway) DeleteMessage(ctx context.Context, request *api.Event) (response *api.Event, err error) {
	message, err := event{request}.validate()
	if err != nil {
		gateway.log.Error("save message error", "err", err)
		return
	}

	err = gateway.Event.Delete(ctx, message)
	if err != nil {
		gateway.log.Error("save message error", "err", err)
		return
	}

	return
}
