package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mendesbarreto/go-my-coffe-shop/internal/user/repository"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/model"
	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
)

func (u *UserGRPCHandler) GetMe(ctx context.Context, req *gen.EmptyRequest) (*gen.GetUserDetailsResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userCached := ctx.Value("user").(model.User)

	user, err := repository.GetUserById(ctx, userCached.ID)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &gen.GetUserDetailsResponse{
		UserId: user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
	}, nil
}
