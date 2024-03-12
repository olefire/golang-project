package repository

import (
	"backend/internal/models"
	UserService "backend/internal/services/user"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

var _ UserService.Repository = (*UserRepository)(nil)

func NewUserRepository(col *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: col,
	}
}

func (ur *UserRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user *models.User

	query := bson.M{"email": email}
	err := ur.collection.FindOne(ctx, query).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) CreateUser(c context.Context, user *models.User) (string, error) {
	result, err := ur.collection.InsertOne(c, user)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
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

func (ur *UserRepository) GetUser(c context.Context, id string) (*models.User, error) {
	var user models.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = ur.collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return &user, err
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id string) error {
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = ur.collection.DeleteOne(ctx, bson.M{"_id": idHex})
	if err != nil {
		return err
	}

	return err
}
