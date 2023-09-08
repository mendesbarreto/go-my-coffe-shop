package handler

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
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

func (u *UserGRPCHandler) SignIn(context.Context, *gen.SignInRequest) (*gen.SignInResponse, error) {
	return &gen.SignInResponse{Token: "1234"}, nil
}

func (u *UserGRPCHandler) SignUp(context.Context, *gen.SignUpRequest) (*gen.SignUpResponse, error) {
}
