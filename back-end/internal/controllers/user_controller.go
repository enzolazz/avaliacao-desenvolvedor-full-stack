package controllers

import (
	"net/http"
	"time"
	"url-shortener/back-end/config"
	"url-shortener/back-end/internal/dtos"
	"url-shortener/back-end/internal/models"
	"url-shortener/back-end/internal/services"
	"url-shortener/back-end/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
	AuthService *services.AuthService
}

func NewUserController(userService *services.UserService, authService *services.AuthService) *UserController {
	return &UserController{UserService: userService, AuthService: authService}
}

func (c *UserController) Register(ctx *gin.Context) {
	var input dtos.RegisterRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	user := models.NewUser(input.Name, input.Surname, input.Username, input.Password)

	if err := c.UserService.CreateUser(user); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, loggedUser, err := c.AuthService.Login(input.Username, input.Password)
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
			ID:       loggedUser.ID.Hex(),
			Name:     loggedUser.Name,
			Surname:  loggedUser.Surname,
			Username: loggedUser.Username,
		},
	})
}

func (c *UserController) GetAll(ctx *gin.Context) {
	users, err := c.UserService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
