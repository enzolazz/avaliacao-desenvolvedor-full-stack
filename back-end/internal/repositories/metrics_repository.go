package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/models"
)

type MetricsRepository struct {
	Collection *mongo.Collection
}

func NewMetricsRepository(col *mongo.Collection) *MetricsRepository {
	return &MetricsRepository{Collection: col}
}

func (r *MetricsRepository) Insert(metric *models.RedirectMetric) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.Collection.InsertOne(ctx, metric)
	return err
}
