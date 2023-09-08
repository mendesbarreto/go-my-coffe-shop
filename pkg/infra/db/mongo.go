package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"

	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
)

var (
	mongoClient  *mongo.Client
	databaseName string
)

func New(ctx context.Context, config *config.Config) *mongo.Client {
	if mongoClient != nil {
		return mongoClient
	}

	opt := options.Client().ApplyURI(config.MongoDb.URI)

	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		slog.Error("Problem to connect with mongodb", "URI", config.MongoDb.URI)
		panic(err)
	}

	slog.Info("Testing Mongo Connection")
	if err = client.Ping(ctx, nil); err != nil {
		slog.Error("Problem to connect with mongodb", "URI", config.MongoDb.URI)
		panic(err)
	}

	databaseName = config.Name

	slog.Info("Mongo Client connection established at", "URI", config.MongoDb.URI)
	mongoClient = client

	return client
}

func GetDatabase() *mongo.Database {
	return mongoClient.Database(databaseName)
}

func Disconnect(ctx context.Context) {
	if err := mongoClient.Disconnect(ctx); err != nil {
		slog.Error("Problem to disconnect mongo db")
	}
}
