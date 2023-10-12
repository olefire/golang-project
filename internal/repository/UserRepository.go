package repository

import (
	"context"
	"fmt"
	"go-testDB/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	Create(c context.Context, user *models.User) error
	Fetch(c context.Context) ([]models.User, error)
	GetByEmail(c context.Context, email string) (models.User, error)
	GetByID(c context.Context, id string) (models.User, error)
}

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *models.User) error {
	collection := ur.database.Collection(ur.collection)

	insertResult, err := collection.InsertOne(c, user)

	if err == nil {
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	}

	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]models.User, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []models.User

	err = cursor.All(c, &users)
	if users == nil {
		return []models.User{}, err
	}

	return users, err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (models.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user models.User

	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (models.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user models.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}
