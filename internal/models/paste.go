package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Paste struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title" binding:"required"`
	Paste  string             `json:"paste" bson:"paste" binding:"required"`
	UserID primitive.ObjectID `json:"userID,omitempty" bson:"userID,omitempty"`
}
