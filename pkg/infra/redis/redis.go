package redis

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var client *redis.Client

func InitRedisClient(uri string) *redis.Client {
	if client != nil {
		slog.Warn("redis", "Redis was already initialized, returning an existent instance")
		return client
	}

	client = redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: "",
		DB:       0,
	})

	return client
}

func GetRedisClient() *redis.Client {
	if client == nil {
		panic("Redis Instance was not initialized")
	}

	return client
}

func Save(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = GetRedisClient().Set(ctx, key, b, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func HasKey(ctx context.Context, key string) bool {
	b, err := GetRedisClient().Get(ctx, key).Result()
	if err != nil {
		return false
	}

	if len(b) == 0 {
		return false
	}

	return true
}

func Get(ctx context.Context, key string, value interface{}) error {
	b, err := GetRedisClient().Get(ctx, key).Result()
	if err != nil {
		return err
	}

	if len(b) == 0 {
		return status.Error(codes.Internal, "Redis key was not found")
	}

	err = json.Unmarshal([]byte(b), value)

	if err != nil {
		return err
	}

	return nil
}

func Disconnect() {
	if client == nil {
		slog.Warn("redis", "No redis client found to disconect")
		return
	}

	err := client.Close()
	if err != nil {
		slog.Error("redis", err.Error())
	}
}
