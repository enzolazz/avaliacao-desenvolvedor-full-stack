package services

import (
	"time"

	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/models"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/repositories"
)

type MetricsService struct {
	Repo *repositories.MetricsRepository
}

func NewMetricsService(repo *repositories.MetricsRepository) *MetricsService {
	return &MetricsService{Repo: repo}
}

func (s *MetricsService) TrackRedirect(shortLinkID string) error {
	metric := &models.RedirectMetric{
		ShortLinkID:  shortLinkID,
		RedirectedAt: time.Now(),
	}

	return s.Repo.Insert(metric)
}
