package handler

import (
	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/product/config"
	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ProductGRPCHandler struct {
	gen.UnimplementedProductServiceServer
	*config.Config
}

var productGRPCHandler *ProductGRPCHandler

func NewProductGRPCHandler(grpcServer *grpc.Server, config *config.Config) *ProductGRPCHandler {
	if productGRPCHandler != nil {
		return productGRPCHandler
	}

	productServiceServer := &ProductGRPCHandler{Config: config}

	gen.RegisterProductServiceServer(grpcServer, productServiceServer)

	if config.EnableGRPCReflection {
		reflection.Register(grpcServer)
	}

	return productServiceServer
}
