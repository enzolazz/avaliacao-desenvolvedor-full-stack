package controllers

import (
	"net/http"

	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/dtos"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/models"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

func (c *UserController) Register(ctx *gin.Context) {
	var input dtos.RegisterRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	user := models.NewUser(input.Name, input.Surname, input.Username, input.Password)

	if err := c.Service.CreateUser(user); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dtos.RegisterResponse{
		ID:       user.ID,
		Name:     user.Name,
		Surname:  user.Surname,
		Username: user.Username,
	})
}

func (c *UserController) GetAll(ctx *gin.Context) {
	users, err := c.Service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	user, err := c.Service.GetUserByID(id)
	if err != nil || user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
