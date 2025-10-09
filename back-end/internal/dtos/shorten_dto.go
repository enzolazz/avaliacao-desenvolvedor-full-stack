package dtos

type CreateShortLinkRequest struct {
	OriginalURL string `json:"original_url"`
}

type CreateShortLinkResponse struct {
	ID          string `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}
