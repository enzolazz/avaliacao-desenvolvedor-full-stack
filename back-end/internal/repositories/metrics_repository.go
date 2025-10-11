package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"url-shortener/back-end/internal/models"
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

func (r *MetricsRepository) aggregateRedirects(shortLinkID string, from, to time.Time, interval string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	from = from.UTC()
	to = to.UTC()

	var unit string
	var step time.Duration
	var binSize int32 = 1

	switch interval {
	case "5min":
		unit = "minute"
		binSize = 5
		step = 5 * time.Minute
		from = from.Truncate(step)
	case "hour":
		unit = "hour"
		step = time.Hour
		from = from.Truncate(step)
	case "day":
		unit = "day"
		step = 24 * time.Hour
		from = from.Truncate(step)
	default:
		unit = "day"
		step = 24 * time.Hour
	}

	pipeline := mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "short_link_id", Value: shortLinkID},
				{Key: "redirected_at", Value: bson.D{
					{Key: "$gte", Value: from},
					{Key: "$lte", Value: to},
				}},
			}},
		},
		bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{
					{Key: "$dateTrunc", Value: bson.D{
						{Key: "date", Value: "$redirected_at"},
						{Key: "unit", Value: unit},
						{Key: "binSize", Value: binSize},
						{Key: "timezone", Value: "UTC"},
					}},
				}},
				{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
			}},
		},
		bson.D{
			{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}},
		},
	}

	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	type resultType struct {
		ID    time.Time `bson:"_id"`
		Count int64     `bson:"count"`
	}

	results := []resultType{}
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	var response []map[string]interface{}
	buckets := make(map[string]int64)
	for _, r := range results {
		buckets[r.ID.Format(time.RFC3339)] = r.Count
	}

	for t := from; !t.After(to); t = t.Add(step) {
		key := t.Format(time.RFC3339)
		response = append(response, map[string]interface{}{
			"time":  key,
			"count": buckets[key],
		})
	}

	return response, nil
}

func (r *MetricsRepository) CountLastHour(shortLinkID string) ([]map[string]any, error) {
	now := time.Now().UTC()
	from := now.Add(-1 * time.Hour)
	return r.aggregateRedirects(shortLinkID, from, now, "5min")
}

func (r *MetricsRepository) CountLastDay(shortLinkID string) ([]map[string]any, error) {
	now := time.Now().UTC()
	from := now.Add(-24 * time.Hour)
	return r.aggregateRedirects(shortLinkID, from, now, "hour")
}

func (r *MetricsRepository) CountLastMonth(shortLinkID string) ([]map[string]any, error) {
	now := time.Now().UTC()
	from := now.AddDate(0, -1, 0)
	return r.aggregateRedirects(shortLinkID, from, now, "day")
}
