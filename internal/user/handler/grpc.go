package handler

import (
	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserGRPCHandler struct {
	gen.UnimplementedUserServiceServer
	*config.Config
}

var userGRPCHandler *UserGRPCHandler

func NewUserGRPCHandler(grpcServer *grpc.Server, config *config.Config) *UserGRPCHandler {
	if userGRPCHandler != nil {
		return userGRPCHandler
	}

	userServiceServer := &UserGRPCHandler{Config: config}

	gen.RegisterUserServiceServer(grpcServer, userServiceServer)

	if config.EnableGRPCReflection {
		reflection.Register(grpcServer)
	}

	return userServiceServer
}
