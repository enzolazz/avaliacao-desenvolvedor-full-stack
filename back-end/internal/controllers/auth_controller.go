package controllers

import (
	"net/http"
	"time"
	"url-shortener/back-end/config"
	"url-shortener/back-end/internal/dtos"
	"url-shortener/back-end/internal/services"
	"url-shortener/back-end/internal/utils"

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

	accessToken, refreshToken, user, err := c.Service.Login(input.Username, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	secure, domain := utils.GetDomain()

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(config.Consts.AccessTokenExp),
		Path:     "/",
		Domain:   domain,
		Secure:   secure,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(config.Consts.RefreshTokenExp),
		Path:     "/",
		Domain:   domain,
		Secure:   secure,
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
	oldRefreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No token provided"})
		return
	}

	accessToken, refreshToken, err := c.Service.RotateRefreshToken(oldRefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	secure, domain := utils.GetDomain()

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(config.Consts.AccessTokenExp),
		Path:     "/",
		Domain:   domain,
		Secure:   secure,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(config.Consts.RefreshTokenExp),
		Path:     "/",
		Domain:   domain,
		Secure:   secure,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	ctx.JSON(http.StatusOK, gin.H{"message": "new tokens issued"})
}

func (c *AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", true, true)

	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", true, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logout realizado com sucesso",
	})
}
