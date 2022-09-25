package pkg

import (
	"context"

	commons "github.com/cryptnode-software/commons/pkg"
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
	log            commons.Logger
	MessageService MessageService
	UserService    UserService
}

func (gateway *Gateway) SaveMessage(ctx context.Context, request *api.Message) (response *api.Message, err error) {

	message, err := message{request}.validate()
	if err != nil {
		gateway.log.Error("save message error", "err", err)
		return
	}

	message, err = gateway.MessageService.Save(ctx, message)
	if err != nil {
		gateway.log.Error("save message error", "err", err)
		return
	}

	response, err = cmessage{message}.convert()
	if err != nil {
		gateway.log.Error("save message error", "err", err)
	}

	return
}

func (gateway *Gateway) DeleteMessage(ctx context.Context, request *api.Message) (response *api.Message, err error) {
	message, err := message{request}.validate()
	if err != nil {
		gateway.log.Error("save message error", "err", err)
		return
	}

	err = gateway.MessageService.Delete(ctx, message)
	if err != nil {
		gateway.log.Error("save message error", "err", err)
		return
	}

	return
}
