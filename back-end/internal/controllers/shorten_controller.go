package controllers

import (
	"net/http"
	"url-shortener/back-end/internal/dtos"
	"url-shortener/back-end/internal/services"
	"url-shortener/back-end/internal/utils"

	"github.com/gin-gonic/gin"
)

type ShortLinkController struct {
	Service *services.ShortLinkService
}

func NewShortLinkController(service *services.ShortLinkService) *ShortLinkController {
	return &ShortLinkController{Service: service}
}

func (c *ShortLinkController) Create(ctx *gin.Context) {
	var input dtos.CreateShortLinkRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return
	}

	alive := utils.IsURLAlive(input.OriginalURL)
	if !alive {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "URL inválida ou inativa"})
		return
	}

	shortLink, err := c.Service.CreateShortLink(userID, input.OriginalURL, input.Label)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dtos.CreateShortLinkResponse{
		ID: shortLink.ID,
	})
}

func (c *ShortLinkController) GetAllByUser(ctx *gin.Context) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return
	}

	links, err := c.Service.GetAllUserLinks(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, links)
}
