package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
)

func (u *UserGRPCHandler) GetMe(ctx context.Context, req *gen.EmptyRequest) (*gen.GetUserDetailsResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return nil, status.Error(codes.NotFound, "This method was not implemented yet")
}
