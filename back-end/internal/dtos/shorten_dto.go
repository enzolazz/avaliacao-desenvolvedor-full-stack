package dtos

type CreateShortLinkRequest struct {
	Label       string `json:"label"`
	OriginalURL string `json:"url"`
}

type CreateShortLinkResponse struct {
	ID string `json:"id"`
}
