package socket

import (
	"github.com/cryptnode-software/hermes/model"
	"github.com/cryptnode-software/hermes/pkg"
)

type sockets map[string][]pkg.Socket

var (
	open = sockets{}
)

type Service struct {
}

func (service *Service) Connect(resource string, socket pkg.Socket) error {
	_, ok := open[resource]
	if !ok {
		open[resource] = []pkg.Socket{}
	}

	open[resource] = append(open[resource], socket)

	return nil
}

func (service *Service) Dispatch(resource string, events []*model.Event) (err error) {
	if sockets, ok := open[resource]; ok {
		for i := range sockets {
			if err = sockets[i].Send(events); err != nil {
				return err
			}
		}
	}

	return nil
}
