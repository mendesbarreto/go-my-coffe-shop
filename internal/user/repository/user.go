package repository

import (
	"context"
	"log/slog"

	"github.com/mendesbarreto/go-my-coffe-shop/pkg/infra/db"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetUserCollection() *mongo.Collection {
	return db.GetDatabase().Collection("user")
}

func GetUserById(ctx context.Context, userId string) (*model.User, error) {
	slog.Info("ID>>>>>>>>>>>>>>>> %v", userId)
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	user := &model.User{}
	err = GetUserCollection().FindOne(ctx, bson.M{"_id": objectId}).Decode(user)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return user, nil
}

func GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	err := GetUserCollection().FindOne(ctx, bson.M{"email": email}).Decode(user)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return user, nil
}