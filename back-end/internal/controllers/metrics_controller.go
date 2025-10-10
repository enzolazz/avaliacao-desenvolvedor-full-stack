package controllers

import (
	"net/http"

	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type MetricsController struct {
	Service *services.MetricsService
}

func NewMetricsController(metricsService *services.MetricsService) *MetricsController {
	return &MetricsController{
		Service: metricsService,
	}
}

func (c *MetricsController) LastHour(ctx *gin.Context) {
	shortLinkID := ctx.Param("id")
	if shortLinkID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	data, err := c.Service.CountLastHour(shortLinkID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *MetricsController) LastDay(ctx *gin.Context) {
	shortLinkID := ctx.Param("id")
	if shortLinkID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	data, err := c.Service.CountLastDay(shortLinkID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *MetricsController) LastMonth(ctx *gin.Context) {
	shortLinkID := ctx.Param("id")
	if shortLinkID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	data, err := c.Service.CountLastMonth(shortLinkID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
