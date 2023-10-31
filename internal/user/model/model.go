package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password"`
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
}
