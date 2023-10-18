package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Paste struct {
	gorm.Model
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title" binding:"required"`
	Paste  string             `json:"paste" bson:"paste" binding:"required"`
	UserID primitive.ObjectID `json:"userID,omitempty" bson:"userID,omitempty"`
}
