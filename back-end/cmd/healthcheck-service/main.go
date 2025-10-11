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

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env from project root!")
	}

	cfg := config.GetConfig()

	db.ConnectMongoDB(cfg.MongoURI, cfg.DBName)
	linksCollection := db.Client.Database(cfg.DBName).Collection("shortlinks")

	ps := pubsub.NewRedisPubSub(cfg)

	ticker := time.NewTicker(config.GetConstants().HealthCheckInterval)
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

		sem := make(chan struct{}, config.GetConstants().MaxGoRoutines)
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
		if inactiveCount >= config.GetConstants().MaxInactiveFailures {
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
