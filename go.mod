module github.com/mendesbarreto/go-my-coffe-shop

go 1.20

require (
	github.com/caarlos0/env/v9 v9.0.0
	github.com/envoyproxy/protoc-gen-validate v1.0.2
	github.com/felixge/httpsnoop v1.0.4
	github.com/golang-jwt/jwt/v5 v5.1.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.0
	github.com/redis/go-redis/v9 v9.7.0
	go.mongodb.org/mongo-driver v1.12.1
	go.uber.org/automaxprocs v1.5.3
	golang.org/x/crypto v0.15.0
	google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/stretchr/testify v1.8.3 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231030173426-d783a09b4405 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.59.0
