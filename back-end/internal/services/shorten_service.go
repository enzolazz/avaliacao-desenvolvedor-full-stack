package services

import (
	"errors"
	"fmt"

	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/models"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/repositories"
	"github.com/teris-io/shortid"
)

type ShortLinkService struct {
	Repo *repositories.ShortLinkRepository
}

func NewShortLinkService(repo *repositories.ShortLinkRepository) *ShortLinkService {
	return &ShortLinkService{Repo: repo}
}

func (s *ShortLinkService) CreateShortLink(originalURL, baseURL string) (*models.ShortLink, error) {
	if originalURL == "" {
		return nil, errors.New("original URL cannot be empty")
	}

	id, _ := shortid.Generate()

	shortLink := &models.ShortLink{
		ID:          id,
		OriginalURL: originalURL,
		ShortURL:    fmt.Sprintf("%s/s/%s", baseURL, id),
	}

	if err := s.Repo.Create(shortLink); err != nil {
		return nil, err
	}

	return shortLink, nil
}

func (s *ShortLinkService) GetByID(id string) (*models.ShortLink, error) {
	return s.Repo.GetByID(id)
}
