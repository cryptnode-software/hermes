package main

import (
	"flag"
	"net/http"
	"os"

	commons "github.com/cryptnode-software/commons/pkg"
	micro "github.com/cryptnode-software/grpc"
	"github.com/cryptnode-software/hermes/gorm"
	"github.com/cryptnode-software/hermes/pkg"
	"github.com/cryptnode-software/hermes/services/event"
	"github.com/cryptnode-software/hermes/services/socket"
	"github.com/cryptnode-software/hermes/services/user"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	api "go.buf.build/grpc/go/thenewlebowski/hermes/v1"
)

const (
	db  = "DB_CONNECTION"
	env = "ENV"
)

func main() {

	socketport := flag.String("socket-port", ":5082", "web socket port")
	httpport := flag.String("grpc-web-port", ":5080", "grpc-web port")
	grpcport := flag.String("grpc-port", ":5081", "grpc port")

	flag.Parse()

	if err := new(gorm.Database).Set(os.Getenv(db)); err != nil {
		panic(err)
	}

	logger := commons.NewLogger(env)

	gw, err := NewGateway()
	if err != nil {
		panic(err)
	}

	logger.Info("starting container...")

	server, err := micro.NewServer()
	if err != nil {
		panic(err)
	}

	grpcserver, err := micro.NewGRPCServer(*grpcport)
	if err != nil {
		panic(err)
	}

	api.RegisterHermesServer(grpcserver.Server(), gw)

	webserver := grpcweb.WrapServer(grpcserver.Server(),
		grpcweb.WithOriginFunc(func(str string) bool {
			return true // change this
		}),
	)

	httpServer := &micro.HttpServer{
		Server: &http.Server{
			Addr: *httpport,
			Handler: http.HandlerFunc(
				func(resp http.ResponseWriter, req *http.Request) {
					webserver.ServeHTTP(resp, req)
				},
			),
		},
	}

	socketserver, err := micro.NewSocketServer(*socketport)
	if err != nil {
		panic(err)
	}

	socket, err := NewSocketServer()
	if err != nil {
		panic(err)
	}

	socketserver.Register(socket.Server())

	server.Add(httpServer)
	server.Add(grpcserver)
	server.Add(socketserver)

	if err = server.Run(); err != nil {
		panic(err)
	}
}

func NewSocketServer() (*socket.Service, error) {
	socket, err := socket.NewService()
	if err != nil {
		return nil, err
	}

	return socket, nil
}

func NewGateway() (gw *pkg.Gateway, err error) {
	gw, err = pkg.NewGateway()
	if err != nil {
		panic(err)
	}

	event, err := event.NewService()
	if err != nil {
		panic(err)
	}
	err = gw.SetEvent(event)
	if err != nil {
		return
	}

	user, err := user.NewService()
	if err != nil {
		panic(err)
	}
	err = gw.SetUser(user)
	if err != nil {
		return
	}

	return
}
