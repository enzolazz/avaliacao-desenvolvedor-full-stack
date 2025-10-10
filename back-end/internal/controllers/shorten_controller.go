package controllers

import (
	"net/http"

	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/dtos"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	userIDValue, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userIDHex, ok := userIDValue.(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id"})
		return
	}

	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id format"})
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
	userIDValue, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userIDHex, ok := userIDValue.(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id"})
		return
	}

	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id format"})
		return
	}

	links, err := c.Service.GetAllUserLinks(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, links)
}

func (c *ShortLinkController) Redirect(ctx *gin.Context) {
	id := ctx.Param("id")

	link, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "short link not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"url": link.OriginalURL,
	})
}
