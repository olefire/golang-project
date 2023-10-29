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

func (ur *PasteRepository) GetPasteByTitle(ctx context.Context, title string) (*models.Paste, error) {
	var paste models.Paste

	err := ur.collection.FindOne(ctx, bson.M{"title": title}).Decode(&paste)
	return &paste, err
}
