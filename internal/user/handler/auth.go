package handler

import (
	"context"
	"time"

	"github.com/mendesbarreto/go-my-coffe-shop/internal/user/repository"
	"github.com/mendesbarreto/go-my-coffe-shop/internal/user/util"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/infra/db"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/infra/redis"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/model"
	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const tokenExpirationTimeHours time.Duration = 48 * time.Hour

func (u *UserGRPCHandler) SignIn(ctx context.Context, req *gen.SignInRequest) (*gen.SignInResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := repository.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword())); err != nil {
		return nil, status.Error(codes.PermissionDenied, "The user or the password does not match the database")
	}

	tokenString, err := util.GenerateJwt(user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = redis.Save(ctx, tokenString, user, tokenExpirationTimeHours)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.SignInResponse{Token: tokenString}, nil
}

func (u *UserGRPCHandler) SignUp(ctx context.Context, req *gen.SignUpRequest) (*gen.SignUpResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var user *model.User = nil
	userCollection := db.GetDatabase().Collection("user")
	if userCollection == nil {
		return nil, status.Error(codes.Internal, "Collection user was not found on the database")
	}

	err = userCollection.FindOne(ctx, bson.M{"email": req.GetEmail()}).Decode(user)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	if user != nil {
		return nil, status.Error(codes.AlreadyExists, "An user already Exists with this email")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "Problem to generate user password")
	}

	user = &model.User{
		ID:       primitive.NewObjectID(),
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: string(passwordHash),
	}

	_, err = userCollection.InsertOne(ctx, *user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	tokenString, err := util.GenerateJwt(user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = redis.Save(ctx, tokenString, user, tokenExpirationTimeHours)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.SignUpResponse{UserId: user.ID.String(), Token: tokenString}, nil
}
