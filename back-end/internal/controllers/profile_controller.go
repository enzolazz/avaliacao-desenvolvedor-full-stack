package controllers

import (
	"net/http"

	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/dtos"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProfileController struct {
	UserService *services.UserService
}

func NewProfileController(userService *services.UserService) *ProfileController {
	return &ProfileController{UserService: userService}
}

func (c *ProfileController) Me(ctx *gin.Context) {
	userIDStr, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in token"})
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	user, err := c.UserService.GetUserByID(userID)
	if err != nil || user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, dtos.UserInfoResponse{
		ID:       user.ID.String(),
		Name:     user.Name,
		Surname:  user.Surname,
		Username: user.Username,
	})
}
