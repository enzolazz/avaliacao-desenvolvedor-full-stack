package routes

import (
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/controllers"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/db"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/repositories"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, jwtSecret string) {
	userRepo := repositories.NewUserRepository(db.Database.Collection("users"))
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", userController.RegisterHandler)
			auth.POST("/login", func(c *gin.Context) { userController.LoginHandler(c, jwtSecret) })
			// auth.GET("/me", controllers.Me)
		}

		usersGroup := api.Group("/users")
		usersGroup.Use(func(c *gin.Context) {
			clientIP := c.ClientIP()
			if clientIP != "127.0.0.1" && clientIP != "::1" {
				c.JSON(403, gin.H{"error": "Forbidden"})
				c.Abort()
				return
			}
			c.Next()
		})
		{
			usersGroup.GET("", userController.GetAllUsersHandler)
			usersGroup.GET("/:id", userController.GetUserByIDHandler)
		}
	}
}
