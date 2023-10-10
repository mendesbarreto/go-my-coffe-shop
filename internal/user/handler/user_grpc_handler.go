package handler

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

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
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password"`
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
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

func (u *UserGRPCHandler) SignIn(ctx context.Context, req *gen.SignInRequest) (*gen.SignInResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user := &User{}
	userCollection := db.GetDatabase().Collection("user")

	err = userCollection.FindOne(ctx, bson.M{"email": req.GetEmail()}).Decode(user)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword())); err != nil {
		return nil, status.Error(codes.PermissionDenied, "The user or the password does not match the database")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user": User{ID: user.ID, Name: user.Name, Email: user.Email}})

	tokenString, err := token.SignedString([]byte(config.GetConfig().AuthSecrete))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.SignInResponse{Token: tokenString}, nil
}

func (u *UserGRPCHandler) GetMe(ctx context.Context, req *gen.EmptyRequest) (*gen.GetUserDetailsResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return nil, status.Error(codes.NotFound, "This method was not implemented yet")
}

func (u *UserGRPCHandler) SignUp(ctx context.Context, req *gen.SignUpRequest) (*gen.SignUpResponse, error) {
	err := req.ValidateAll()
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
