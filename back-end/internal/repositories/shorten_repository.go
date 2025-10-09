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
