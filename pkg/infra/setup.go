package infra

import (
	"context"

	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/infra/db"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/infra/redis"
)

func SetupDependecies(ctx context.Context, config *config.Config) {
	db.New(ctx, config)
	redis.InitRedisClient(config.Redis.URI)
}

func CleanUpDependcies(ctx context.Context) {
	db.Disconnect(ctx)
	redis.Disconnect()
}
