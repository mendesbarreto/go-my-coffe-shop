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

var productServiceClient gen.ProductServiceClient

func GetProductServiceClient(host string, port string) (gen.ProductServiceClient, error) {
	if productServiceClient != nil {
		return productServiceClient, nil
	}
	// The WithInsecure enables the developer to connect localhost or http
	// The WithBlock block any call to the server until the connectio is up
	serverAddress := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.DialContext(context.Background(), serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	if conn == nil {
		slog.Error("Problem to connect ", config.GetConfig().UserServiceHost)
		return nil, status.Error(codes.Unavailable, "Service user not available")
	}

	productServiceClient = gen.NewProductServiceClient(conn)

	slog.Info("User client connected", conn.GetState())
	return productServiceClient, nil
}
