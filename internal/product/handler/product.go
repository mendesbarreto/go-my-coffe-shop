package handler

import (
	"context"

	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
)

func (h *ProductGRPCHandler) ListProducts(context.Context, *gen.ListProductsRequest) (*gen.ListProductsResponse, error) {
	products := []*gen.ProductDetails{
		{
			Id:    "1",
			Name:  "Strawberry Cake",
			Price: 10.0,
		},
		{
			Id:    "2",
			Name:  "Chocolate Cake",
			Price: 15.0,
		},
		{
			Id:    "3",
			Name:  "Brigadeiro",
			Price: 20.0,
		},

		{
			Id:    "3",
			Name:  "Coffe",
			Price: 3.0,
		},
	}

	return &gen.ListProductsResponse{
		Items: products,
	}, nil
}
