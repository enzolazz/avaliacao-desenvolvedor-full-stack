package dtos

import "time"

type CreateShortLinkRequest struct {
	Label       string `json:"label"`
	OriginalURL string `json:"url"`
}

type CreateShortLinkResponse struct {
	ID string `json:"id"`
}

type GetShortLinkResponse struct {
	ID          string    `bson:"_id" json:"id"`
	Label       string    `bson:"label" json:"label"`
	OriginalURL string    `bson:"original_url" json:"original_url"`
	Status      string    `bson:"status" json:"status"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
}
