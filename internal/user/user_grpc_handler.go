package handler

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log/slog"

	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/infra/db"
	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
)

type UserGRPCHandler struct {
	gen.UnimplementedUserServiceServer
	*config.Config
}

var userGRPCHandler *UserGRPCHandler

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitepty"`
	Email    string             `bson: "email,omitempty"`
	Password string             `bson: "password"`
}

func NewUserGRPCHandler(grpcServer *grpc.Server, config *config.Config) *UserGRPCHandler {
	if userGRPCHandler != nil {
		return userGRPCHandler
	}

	userServiceServer := &UserGRPCHandler{Config: config}

	gen.RegisterUserServiceServer(grpcServer, userServiceServer)

	if config.EnableGRPCReflection {
		reflection.Register(grpcServer)
	}

	return userServiceServer
}

func (u *UserGRPCHandler) SignIn(context.Context, *gen.SignInRequest) (*gen.SignInResponse, error) {
	return &gen.SignInResponse{Token: "1234"}, nil
}

func (u *UserGRPCHandler) SignUp(ctx context.Context, req *gen.SignUpRequest) (*gen.SignUpResponse, error) {
	slog.Info("Validatin request data")
	err := req.ValidateAll()
	slog.Info("Request data validated")
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var user *User
	userCollection := db.GetDatabase().Collection("user")
	if userCollection == nil {
		return nil, status.Error(codes.Internal, "Collection user was not found on the database")
	}

	err = userCollection.FindOne(ctx, bson.M{"email": req.GetEmail()}).Decode(&user)

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

	user = &User{ID: primitive.NewObjectID(), Name: req.GetName(), Email: req.GetEmail(), Password: string(passwordHash)}

	_, err = userCollection.InsertOne(ctx, *user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.SignUpResponse{UserId: user.ID.String(), Token: "123"}, nil
}
