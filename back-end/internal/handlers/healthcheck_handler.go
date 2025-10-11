package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"url-shortener/back-end/internal/models"
	"url-shortener/back-end/internal/pubsub"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LinkStatusUpdate struct {
	UserID        string    `json:"userId"`
	ShortLinkID   string    `json:"shortLinkId"`
	Status        string    `json:"status"`
	InactiveCount int       `json:"inactiveCount"`
	LastCheckedAt time.Time `json:"lastCheckedAt"`
}

type NotifyFunc func(update LinkStatusUpdate)

func HandleLinkStatusUpdates(ps *pubsub.RedisPubSub, coll *mongo.Collection) (func(), error) {
	unsubscribe, err := ps.Subscribe("links:status", func(msg string) {
		var update struct {
			ShortLinkID string `json:"shortLinkId"`
			Status      string `json:"status"`
		}

		if err := json.Unmarshal([]byte(msg), &update); err != nil {
			log.Println("Failed to decode healthcheck message:", err)
			return
		}

		var link models.ShortLink
		if err := coll.FindOne(context.Background(), bson.M{"_id": update.ShortLinkID}).Decode(&link); err != nil {
			log.Println("Failed to find short link:", err)
			return
		}

		userChannel := fmt.Sprintf("user:%s:updates", link.UserID.Hex())
		userMsg, _ := json.Marshal(update)

		log.Println("Trying to send updates to user: " + link.UserID.Hex())
		if err := ps.Publish(userChannel, string(userMsg)); err != nil {
			log.Println("Failed to publish user-specific update:", err)
		}
	})
	if err != nil {
		return nil, err
	}

	log.Println("LinkStatusHandler started successfully")
	return unsubscribe, nil
}
