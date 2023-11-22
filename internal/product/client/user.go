package client

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/product/config"
	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var userServiceClient gen.UserServiceClient

func GetUserServiceClient() (gen.UserServiceClient, error) {
	if userServiceClient != nil {
		return userServiceClient, nil
	}
	// The WithInsecure enables the developer to connect localhost or http
	// The WithBlock block any call to the server until the connectio is up
	serverAddress := fmt.Sprintf("%s:%s", config.GetConfig().UserServiceHost, config.GetConfig().UserServicePort)
	conn, err := grpc.DialContext(context.Background(), serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	if conn == nil {
		slog.Error("Problem to connect ", config.GetConfig().UserServiceHost)
		return nil, status.Error(codes.Unavailable, "Service user not available")
	}

	userServiceClient = gen.NewUserServiceClient(conn)

	slog.Info("User client connected", conn.GetState())

	return userServiceClient, nil
}
