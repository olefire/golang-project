package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name" binding:"required"`
	Email string             `json:"email" bson:"email" binding:"email,required"`
}
