module github.com/mendesbarreto/go-my-coffe-shop

go 1.20

require (
	github.com/caarlos0/env/v9 v9.0.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.17.1
	go.uber.org/automaxprocs v1.5.3
	google.golang.org/genproto/googleapis/api v0.0.0-20230822172742-b8732ec3820d
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/genproto v0.0.0-20230821184602-ccc8af3d0e93 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.57.0
