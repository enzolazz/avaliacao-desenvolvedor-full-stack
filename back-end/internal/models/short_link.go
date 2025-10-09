package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShortLink struct {
	ID            string             `bson:"_id" json:"id"`
	UserID        primitive.ObjectID `bson:"user_id" json:"user_id"`
	Label         string             `bson:"label" json:"label"`
	OriginalURL   string             `bson:"original_url" json:"original_url"`
	Status        string             `bson:"status" json:"status"`
	InactiveCount int                `bson:"inactive_count" json:"inactive_count"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
}

func NewShortLink(userID primitive.ObjectID, id, originalURL, label string) ShortLink {
	return ShortLink{
		ID:            id,
		UserID:        userID,
		OriginalURL:   originalURL,
		Label:         label,
		Status:        "active",
		InactiveCount: 0,
		CreatedAt:     time.Now(),
	}
}
