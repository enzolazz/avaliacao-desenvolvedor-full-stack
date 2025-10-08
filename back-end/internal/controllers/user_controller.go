package controllers

import (
	"net/http"

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

func (c *UserController) RegisterHandler(ctx *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.NewUser(input.Name, input.Surname, input.Username, input.Password)
	if err := c.Service.CreateUser(user); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"name":     user.Name,
		"surname":  user.Surname,
		"username": user.Username,
	})
}

func (c *UserController) GetAllUsersHandler(ctx *gin.Context) {
	users, err := c.Service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUserByIDHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	user, _ := c.Service.GetUserByID(id)
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) LoginHandler(ctx *gin.Context, jwtSecret string) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.Service.LoginHandler(input.Username, input.Password, jwtSecret)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
