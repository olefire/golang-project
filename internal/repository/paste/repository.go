package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mongoGo/internal/models"
	PasteService "mongoGo/internal/services/paste"
)

type PasteRepository struct {
	collection *mongo.Collection
}

var _ PasteService.Repository = (*PasteRepository)(nil)

func NewPasteRepository(col *mongo.Collection) *PasteRepository {
	return &PasteRepository{
		collection: col,
	}
}

func (ur *PasteRepository) CreatePaste(ctx context.Context, paste *models.Paste) (string, error) {
	result, err := ur.collection.InsertOne(ctx, paste)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (ur *PasteRepository) GetBatch(ctx context.Context) ([]models.Paste, error) {
	opts := options.Find()
	cursor, err := ur.collection.Find(ctx, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var batch []models.Paste

	err = cursor.All(ctx, &batch)

	if err != nil {
		return nil, err
	}

	return batch, err
}

func (ur *PasteRepository) GetPasteById(ctx context.Context, id string) (*models.Paste, error) {
	var paste models.Paste

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = ur.collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&paste)
	return &paste, err
}

func (ur *PasteRepository) DeletePaste(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = ur.collection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return err
	}

	return err
}
