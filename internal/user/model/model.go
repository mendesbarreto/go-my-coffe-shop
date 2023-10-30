package model

import (
	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
