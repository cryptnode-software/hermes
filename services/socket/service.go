package socket

import (
	"fmt"

	commons "github.com/cryptnode-software/commons/pkg"
	db "github.com/cryptnode-software/hermes/gorm"
	socketio "github.com/googollee/go-socket.io"
	"gorm.io/gorm"
)

type sockets map[string][]Socket

var (
	open = sockets{}
)

type Service struct {
	db     *gorm.DB
	log    commons.Logger
	server *socketio.Server
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

	service.server = socketio.NewServer(nil)

	service.server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	service.server.OnEvent("/", "msg", func(s socketio.Conn) {
		msg := s.Context().(string)
		s.Emit("msg", msg)
		return
	})

	return
}

func (service Service) Server() *socketio.Server {
	return service.server
}
