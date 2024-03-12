package repository

import (
	"backend/internal/models"
	PasteService "backend/internal/services/paste"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (pr *PasteRepository) CreatePaste(ctx context.Context, paste *models.Paste) (string, error) {
	result, err := pr.collection.InsertOne(ctx, paste)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (pr *PasteRepository) GetBatch(ctx context.Context) ([]models.Paste, error) {
	opts := options.Find()
	cursor, err := pr.collection.Find(ctx, bson.D{}, opts)

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

func (pr *PasteRepository) GetPasteById(ctx context.Context, id string) (*models.Paste, error) {
	var paste models.Paste

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = pr.collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&paste)
	return &paste, err
}

func (pr *PasteRepository) DeletePaste(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = pr.collection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return err
	}

	return err
}

func (pr *PasteRepository) UpdatePaste(ctx context.Context, id string, upd *models.UpdatePaste) (*models.Paste, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{{Key: "$set", Value: upd}}

	res := pr.collection.FindOneAndUpdate(ctx, query, update)

	var updatedPaste *models.Paste

	if err := res.Decode(&updatedPaste); err != nil {
		return nil, errors.New("no paste with that Id exists")
	}

	return updatedPaste, err
}
