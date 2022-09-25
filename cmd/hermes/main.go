package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	commons "github.com/cryptnode-software/commons/pkg"
	"github.com/cryptnode-software/hermes/pkg"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	api "go.buf.build/grpc/go/thenewlebowski/hermes/v1"
	"google.golang.org/grpc"
)

const (
	env = "ENV"
)

func main() {

	port := flag.Int("port", 5080, "grpc port")

	flag.Parse()

	environ := commons.Environment(os.Getenv(env))
	if environ == "" {
		log.Fatalf("environment is not provided: please provide %s variable", env)
		return
	}

	environment := pkg.NewEnv(commons.NewLogger(environ))

	gw, err := pkg.NewGateway(environment)
	if err != nil {
		panic(err)
	}

	logger := environment.Log
	logger.Info("starting container...")

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(

				grpc_recovery.UnaryServerInterceptor(
					grpc_recovery.WithRecoveryHandlerContext(
						func(ctx context.Context, p interface{}) error {
							logger.Error("grpc_recovery", p, ctx)

							return p.(error)
						},
					),
				),
			),
		),
	}

	grpcServer := grpc.NewServer(opts...)
	api.RegisterHermesServer(grpcServer, gw)

	server := grpcweb.WrapServer(grpcServer,
		grpcweb.WithOriginFunc(func(str string) bool {
			return true // change this
		}),
	)

	handler := func(resp http.ResponseWriter, req *http.Request) {
		server.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: http.HandlerFunc(handler),
	}

	logger.Info(fmt.Sprintf("listening on port :%d", *port))
	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
