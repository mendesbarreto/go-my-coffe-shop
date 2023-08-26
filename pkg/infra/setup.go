package infra

import (
	"context"

	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/infra/db"
)

func SetupDependecies(ctx context.Context, config *config.Config) {
	db.New(ctx, config)
}

func CleanUpDependcies(ctx context.Context) {
	db.Disconnect(ctx)
}
