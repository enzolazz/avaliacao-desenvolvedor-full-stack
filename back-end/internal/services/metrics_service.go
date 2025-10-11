package services

import (
	"time"

	"url-shortener/back-end/internal/models"
	"url-shortener/back-end/internal/repositories"
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

func (s *MetricsService) CountLastHour(shortLinkID string) ([]map[string]any, error) {
	return s.Repo.CountLastHour(shortLinkID)
}

func (s *MetricsService) CountLastDay(shortLinkID string) ([]map[string]any, error) {
	return s.Repo.CountLastDay(shortLinkID)
}

func (s *MetricsService) CountLastMonth(shortLinkID string) ([]map[string]any, error) {
	return s.Repo.CountLastMonth(shortLinkID)
}
