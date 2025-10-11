package main

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"url-shortener/back-end/config"
	"url-shortener/back-end/internal/db"
	"url-shortener/back-end/internal/models"
	"url-shortener/back-end/internal/pubsub"
	"url-shortener/back-end/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	db.ConnectMongoDB(config.Cfg.MongoURI, config.Cfg.DBName)
	linksCollection := db.Client.Database(config.Cfg.DBName).Collection("shortlinks")

	ps := pubsub.NewRedisPubSub()

	ticker := time.NewTicker(config.Consts.HealthCheckInterval)
	defer ticker.Stop()

	log.Println("Healthchecker service started")

	for range ticker.C {
		cursor, err := linksCollection.Find(context.Background(), bson.M{"status": "active"})
		if err != nil {
			log.Println("Failed to fetch short links:", err)
			continue
		}

		var links []models.ShortLink
		if err := cursor.All(context.Background(), &links); err != nil {
			log.Println("Failed to decode short links:", err)
			continue
		}

		sem := make(chan struct{}, config.Consts.MaxGoRoutines)
		for _, link := range links {
			sem <- struct{}{}
			go func(l models.ShortLink) {
				defer func() { <-sem }()
				checkAndPublish(ps, linksCollection, l)
			}(link)
		}

		for i := 0; i < cap(sem); i++ {
			sem <- struct{}{}
		}
	}
}

func checkAndPublish(ps *pubsub.RedisPubSub, coll *mongo.Collection, link models.ShortLink) {
	status := "active"
	inactiveCount := link.InactiveCount

	if !utils.IsURLAlive(link.OriginalURL) {
		inactiveCount++
		if inactiveCount >= config.Consts.MaxInactiveFailures {
			status = "inactive"
		}
	} else {
		inactiveCount = 0
	}

	if status != link.Status {
		update := bson.M{
			"$set": bson.M{
				"status":          status,
				"inactive_count":  inactiveCount,
				"last_checked_at": time.Now().UTC(),
			},
		}
		_, err := coll.UpdateByID(context.Background(), link.ID, update)
		if err != nil {
			log.Println("Failed to update link status:", err)
			return
		}

		if status == "inactive" {
			msg, _ := json.Marshal(map[string]any{
				"shortLinkId": link.ID,
				"status":      status,
			})
			log.Println("Trying to publish inactive status to link: " + link.ID)
			if err := ps.Publish("links:status", string(msg)); err != nil {
				log.Println("Failed to publish update:", err)
			}
		}
	} else {
		if inactiveCount != link.InactiveCount {
			_, _ = coll.UpdateByID(context.Background(), link.ID, bson.M{
				"$set": bson.M{"inactive_count": inactiveCount, "last_checked_at": time.Now().UTC()},
			})
		}
	}
}
