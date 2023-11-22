package db

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient  *mongo.Client
	databaseName string
)

type DBConfig interface{}

func New(ctx context.Context, uri string, dbName string) *mongo.Client {
	if mongoClient != nil {
		return mongoClient
	}

	opt := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		slog.Error("Problem to connect with mongodb", "URI", uri)
		panic(err)
	}

	slog.Info("Testing Mongo Connection")
	if err = client.Ping(ctx, nil); err != nil {
		slog.Error("Problem to connect with mongodb", "URI", uri)
		panic(err)
	}

	databaseName = dbName

	slog.Info("Mongo Client connection established at", "URI", uri)
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
