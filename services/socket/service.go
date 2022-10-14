package socket

import (
	"context"

	commons "github.com/cryptnode-software/commons/pkg"
	db "github.com/cryptnode-software/hermes/gorm"
	"gorm.io/gorm"
)

type sockets map[string][]Socket

var (
	open = sockets{}
)

type Service struct {
	db  *gorm.DB
	log commons.Logger
}

type Socket interface {
	Send([]*db.Event) error
}

func NewService() (service *Service, err error) {
	service = new(Service)
	if service.db, err = db.Get(); err != nil {
		return
	}

	if service.log, err = commons.GetLogger(); err != nil {
		return
	}
	return
}

func (service *Service) Connect(ctx context.Context, resource string, socket Socket) error {
	_, ok := open[resource]
	if !ok {
		open[resource] = make([]Socket, 0)
	}

	service.log.Info("user connected to socket", "resource", resource)
	open[resource] = append(open[resource], socket)

	return nil
}

func (service *Service) Dispatch(ctx context.Context, resource string, events []*db.Event) (err error) {
	if sockets, ok := open[resource]; ok {
		for i := range sockets {
			if err = sockets[i].Send(events); err != nil {
				return err
			}
		}
	}

	return nil
}
