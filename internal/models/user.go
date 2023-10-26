package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name" binding:"required"`
	Email string             `json:"email" bson:"email" binding:"email,required"`
}
