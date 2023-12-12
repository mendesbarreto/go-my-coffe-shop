# go-my-coffe-shop

The idea is to test technologies and build a simple backend for a coffee shop.

### Run the project

We need to install the project dependencies

```bash
brew install bufbuild/buf/buf
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

cd proto/
buf mod update
cd ..

buf generate
```

Run

## Tech Stack

- [ ] Make gRPC gateway run [REST](https://github.com/grpc-ecosystem/grpc-gateway) Locally
- [ ] gRPC for internal microservices communication
- [ ] REST for external communication (I will try to pen a graphQL after, let's see)
- [ ] MongoDB for the database
- [ ] Use Passkey for login
- [ ] Use Docker for deployment
- [ ] Use Kubernetes for orchestration
- [ ] Use Prometheus for monitoring
- [ ] Use Grafana for visualization
- [ ] Test Jaeager for tracing
- [ ] Use logz.io for logging
- [ ] Use pulumi for infrastructure as code
- [ ] Use solid-js for the frontend
