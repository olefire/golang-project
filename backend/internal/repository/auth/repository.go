package repository

import (
	"backend/internal/models"
	AuthService "backend/internal/services/auth"
	utils "backend/internal/utils/auth"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	collection *mongo.Collection
}

var _ AuthService.Repository = (*AuthRepository)(nil)

func NewUserRepository(col *mongo.Collection) *AuthRepository {
	return &AuthRepository{
		collection: col,
	}
}

func (ar *AuthRepository) SignUpUser(ctx context.Context, user *models.User) (string, error) {
	user.Password, _ = utils.HashPassword(user.Password)
	result, err := ar.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (ar *AuthRepository) SignInUser(ctx context.Context, signInInput *models.SignInInput) (string, error) {
	//TODO implement me
	panic("implement me")
}
