package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RedirectMetric struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShortLinkID  string             `bson:"short_link_id" json:"short_link_id"`
	RedirectedAt time.Time          `bson:"redirected_at" json:"redirected_at"`
}
