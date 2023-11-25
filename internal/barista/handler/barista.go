package handler

import (
	"context"
	"log/slog"

	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/barista/config"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/client"
	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (b *BaristaGRPCHandler) GetMenu(ctx context.Context, request *gen.GetMenuRequest) (*gen.GetMenuResponse, error) {
	slog.Info("HELOOOO >>>>>>>>>>>")
	config := config.GetConfig()
	service, err := client.GetProductServiceClient(config.ProductServiceHost, config.ProductServicePort)
	if err != nil {
		return nil, err
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.NotFound, "Missing headers")
	}

	outgoingCtx := metadata.NewOutgoingContext(ctx, md)
	result, err := service.ListProducts(outgoingCtx, &gen.ListProductsRequest{Ids: request.Ids})
	if err != nil || result == nil {
		return nil, status.Errorf(codes.Unauthenticated, "Menu unavailable %v", err.Error())
	}

	if err != nil {
		return nil, err
	}

	return &gen.GetMenuResponse{Items: result.Items}, nil
}
