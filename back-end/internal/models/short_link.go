package models

type ShortLink struct {
	ID          string `bson:"_id" json:"id"`
	OriginalURL string `bson:"original_url" json:"original_url"`
	ShortURL    string `bson:"short_url" json:"short_url"`
}
