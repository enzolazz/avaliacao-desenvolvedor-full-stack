package repositories

import (
	"context"
	"time"

	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShortLinkRepository struct {
	Collection *mongo.Collection
}

func NewShortLinkRepository(col *mongo.Collection) *ShortLinkRepository {
	return &ShortLinkRepository{Collection: col}
}

func (r *ShortLinkRepository) Create(link *models.ShortLink) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.Collection.InsertOne(ctx, link)
	return err
}

func (r *ShortLinkRepository) Exists(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	count, err := r.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *ShortLinkRepository) GetByID(id string) (*models.ShortLink, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var link models.ShortLink
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&link)
	if err != nil {
		return nil, err
	}

	return &link, nil
}
