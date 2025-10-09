package controllers

import (
	"net/http"

	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/dtos"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/services"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}

	baseURL := scheme + "://" + ctx.Request.Host

	shortLink, err := c.Service.CreateShortLink(input.OriginalURL, baseURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dtos.CreateShortLinkResponse{
		ID:          shortLink.ID,
		OriginalURL: shortLink.OriginalURL,
		ShortURL:    shortLink.ShortURL,
	})
}

func (c *ShortLinkController) Redirect(ctx *gin.Context) {
	id := ctx.Param("id")

	link, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "short link not found"})
		return
	}

	ctx.Redirect(http.StatusFound, link.OriginalURL)
}
