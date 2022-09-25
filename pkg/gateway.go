package pkg

import (
	"context"

	api "go.buf.build/grpc/go/thenewlebowski/hermes/v1"
)

func NewGateway(environment *Environment) (*Gateway, error) {
	return &Gateway{
		environment,
	}, nil
}

type Gateway struct {
	*Environment
}

func (gateway *Gateway) SaveMessage(ctx context.Context, message *api.Message) (result *api.Message, err error) {
	return
}

func (gateway *Gateway) DeleteMessage(ctx context.Context, message *api.Message) (result *api.Message, err error) {
	return
}
