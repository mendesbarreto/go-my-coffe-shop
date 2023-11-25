package handler

import (
	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/barista/config"
	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type BaristaGRPCHandler struct {
	gen.UnimplementedBaristaServiceServer
	*config.Config
}

var barsitaGRPCHandler *BaristaGRPCHandler

func NewBaristaGRPCHandler(grpcServer *grpc.Server, config *config.Config) *BaristaGRPCHandler {
	if barsitaGRPCHandler != nil {
		return barsitaGRPCHandler
	}

	baristaServiceServer := &BaristaGRPCHandler{Config: config}

	gen.RegisterBaristaServiceServer(grpcServer, baristaServiceServer)

	if config.EnableGRPCReflection {
		reflection.Register(grpcServer)
	}

	return baristaServiceServer
}
