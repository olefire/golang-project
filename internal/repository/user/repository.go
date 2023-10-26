package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mongoGo/internal/models"
	"mongoGo/internal/services"
)

type UserRepository struct {
	collection *mongo.Collection
}

var _ services.UserRepository = (*UserRepository)(nil)

func NewUserRepository(col *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: col,
	}
}

func (ur *UserRepository) CreateUser(c context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	insertedId, err := ur.collection.InsertOne(c, user)

	return insertedId, err
}

func (ur *UserRepository) GetUsers(c context.Context) ([]models.User, error) {
	opts := options.Find()
	cursor, err := ur.collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []models.User

	err = cursor.All(c, &users)

	if err != nil {
		return nil, err
	}

	return users, err
}

func (ur *UserRepository) GetByEmail(c context.Context, email string) (*models.User, error) {
	var user models.User

	err := ur.collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return &user, err
}

func (ur *UserRepository) GetByID(c context.Context, id string) (*models.User, error) {
	var user models.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &user, err
	}

	err = ur.collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return &user, err
}
