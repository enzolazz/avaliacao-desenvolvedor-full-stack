package services

import (
	"errors"
	"fmt"

	"url-shortener/back-end/internal/models"
	"url-shortener/back-end/internal/repositories"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShortLinkService struct {
	Repo *repositories.ShortLinkRepository
}

func NewShortLinkService(repo *repositories.ShortLinkRepository) *ShortLinkService {
	return &ShortLinkService{Repo: repo}
}

func (s *ShortLinkService) CreateShortLink(userID primitive.ObjectID, originalURL, label string) (*models.ShortLink, error) {
	if originalURL == "" {
		return nil, errors.New("original URL cannot be empty")
	}

	var shortLink models.ShortLink
	var exists bool

	for {
		id, err := shortid.Generate()
		if err != nil {
			return nil, fmt.Errorf("failed to generate short ID: %w", err)
		}

		exists, err = s.Repo.Exists(id)
		if err != nil {
			return nil, fmt.Errorf("failed to check short ID uniqueness: %w", err)
		}

		if !exists {
			shortLink = models.NewShortLink(userID, id, originalURL, label)
			break
		}
	}

	if err := s.Repo.Create(&shortLink); err != nil {
		return nil, err
	}

	return &shortLink, nil
}

func (s *ShortLinkService) GetAllUserLinks(userID primitive.ObjectID) ([]models.ShortLink, error) {
	return s.Repo.FindAllByUser(userID)
}

func (s *ShortLinkService) GetByID(id string) (*models.ShortLink, error) {
	return s.Repo.GetByID(id)
}
