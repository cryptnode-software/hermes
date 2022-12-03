package socket

import (
	"fmt"
	"net/http"

	commons "github.com/cryptnode-software/commons/pkg"
	db "github.com/cryptnode-software/hermes/gorm"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"gorm.io/gorm"
)

type sockets map[string][]Socket

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

	service.server =
		socketio.NewServer(&engineio.Options{
			Transports: []transport.Transport{
				&polling.Transport{
					CheckOrigin: func(r *http.Request) bool {
						return true
					},
				},
				&websocket.Transport{
					CheckOrigin: func(r *http.Request) bool {
						return true
					},
				},
			},
		})

	service.server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		s.Join("chat")
		return nil
	})

	service.server.OnEvent("/", "message", func(s socketio.Conn, message string) {
		service.server.BroadcastToRoom("", "chat", "message", message)
		return
	})

	service.server.OnEvent("/", "connect", func(s socketio.Conn, user string) {
		service.server.BroadcastToRoom("", "chat", "connect", user)
	})

	return
}

func (service Service) Server() *socketio.Server {
	return service.server
}
