package controllers

import (
	"fmt"
	"net/http"
	"url-shortener/back-end/internal/services"

	"github.com/gin-gonic/gin"
)

type RedirectController struct {
	ShortLinkService *services.ShortLinkService
	MetricsService   *services.MetricsService
}

func NewRedirectController(slService *services.ShortLinkService, mService *services.MetricsService) *RedirectController {
	return &RedirectController{
		ShortLinkService: slService,
		MetricsService:   mService,
	}
}

func (c *RedirectController) HandleRedirect(ctx *gin.Context) {
	id := ctx.Param("id")

	link, err := c.ShortLinkService.GetByID(id)
	if err != nil || link.Status != "active" {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "short link not found"})
		return
	}

	go func() {
		if err := c.MetricsService.TrackRedirect(link.ID); err != nil {
			fmt.Printf("failed to record redirect metric: %v\n", err)
		}
	}()

	ctx.JSON(http.StatusOK, gin.H{
		"url": link.OriginalURL,
	})
}
