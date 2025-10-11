package controllers

import (
	"net/http"
	"time"
	"url-shortener/back-end/internal/dtos"
	"url-shortener/back-end/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{Service: service}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var input dtos.LoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	token, user, err := c.Service.Login(input.Username, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 1),
		Path:    "/",
		// Domain:   domain, // Uncomment for production
		// Secure:   secure, // Uncomment for production
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	ctx.JSON(http.StatusOK, dtos.LoginResponse{
		User: dtos.UserInfoResponse{
			ID:       user.ID.Hex(),
			Name:     user.Name,
			Surname:  user.Surname,
			Username: user.Username,
		},
	})
}

func (c *AuthController) Refresh(ctx *gin.Context) {
	var input dtos.RefreshRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	newToken, err := c.Service.Refresh(input.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dtos.RefreshResponse{Token: newToken})
}
