module github.com/cryptnode-software/hermes

// replace go.buf.build/grpc/go/thenewlebowski/hermes => ../api/gen/proto/go/hermes/

// replace github.com/cryptnode-software/commons => ../commons

go 1.19

require (
	github.com/cryptnode-software/commons v0.0.0-20220925001958-6c78062b33ae
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/improbable-eng/grpc-web v0.15.0
	go.buf.build/grpc/go/thenewlebowski/hermes v1.4.1
	google.golang.org/grpc v1.49.0
)

require (
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/rs/cors v1.7.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20210126160654-44e461bb6506 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	nhooyr.io/websocket v1.8.6 // indirect
)
