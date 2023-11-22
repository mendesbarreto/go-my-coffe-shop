package infra

import (
	"context"

	"github.com/mendesbarreto/go-my-coffe-shop/pkg/infra/db"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/infra/redis"
)

func SetupDependecies(ctx context.Context, dbName string, mongoURI string, redisURI string) {
	db.New(ctx, mongoURI, dbName)
	redis.InitRedisClient(redisURI)
}

func CleanUpDependcies(ctx context.Context) {
	db.Disconnect(ctx)
	redis.Disconnect()
}
