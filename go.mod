module github.com/cryptnode-software/hermes

// replace go.buf.build/grpc/go/thenewlebowski/hermes => ../cryptnode-commons/gen/proto/go/hermes/

// replace github.com/cryptnode-software/commons => ../commons

// replace github.com/cryptnode-software/grpc => ../grpc

go 1.18

require (
	github.com/cryptnode-software/commons v0.0.0-20220925045416-b3e6361f83e0
	github.com/cryptnode-software/grpc v0.0.0-20221123231124-391714d8a847
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/googollee/go-socket.io v1.6.2
	github.com/improbable-eng/grpc-web v0.15.0
	github.com/kr/pretty v0.1.0
	go.buf.build/grpc/go/thenewlebowski/hermes v1.4.2
	gorm.io/driver/mysql v1.3.6
	gorm.io/gorm v1.23.10
)

require (
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/gomodule/redigo v1.8.4 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/kr/text v0.1.0 // indirect
	github.com/rs/cors v1.7.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210126160654-44e461bb6506 // indirect
	google.golang.org/grpc v1.50.1 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	nhooyr.io/websocket v1.8.6 // indirect
)
