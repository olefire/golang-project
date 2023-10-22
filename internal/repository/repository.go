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
	database   *mongo.Database
	collection string
}

var _ services.Repository = (*UserRepository)(nil)

func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *UserRepository) CreateUser(c context.Context, user *models.User) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *UserRepository) GetUsers(c context.Context) ([]models.User, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find()
	cursor, err := collection.Find(c, bson.D{}, opts)

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
	collection := ur.database.Collection(ur.collection)

	var user *models.User

	err := collection.FindOne(c, bson.M{"email": email}).Decode(user)
	return user, err
}

func (ur *UserRepository) GetByID(c context.Context, id string) (*models.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user *models.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}