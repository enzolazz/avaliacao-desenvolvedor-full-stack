package routes

import (
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/controllers"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/db"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/middleware"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/repositories"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, jwtSecret string) {
	userRepo := repositories.NewUserRepository(db.Database.Collection("users"))
	shortLinkRepo := repositories.NewShortLinkRepository(db.Database.Collection("shortlinks"))

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo, jwtSecret)
	shortLinkService := services.NewShortLinkService(shortLinkRepo)

	userController := controllers.NewUserController(userService)
	profileController := controllers.NewProfileController(userService)
	authController := controllers.NewAuthController(authService)
	shortLinkController := controllers.NewShortLinkController(shortLinkService)

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		auth.POST("/login", authController.Login)
		auth.POST("/register", userController.Register)

		// Protected route
		profile := api.Group("/profile")
		profile.Use(middleware.JWTMiddleware(jwtSecret))
		profile.GET("/me", profileController.Me)

		users := api.Group("/users")
		users.Use(middleware.AllowOnlyLocalhost())
		users.GET("", userController.GetAll)
		users.GET("/:id", userController.GetByID)

		// Protected route
		shortlinks := api.Group("/shorten")
		shortlinks.Use(middleware.JWTMiddleware(jwtSecret))
		shortlinks.POST("", shortLinkController.Create)

	}
	// Public search route
	r.GET("/redirect/:id", shortLinkController.Redirect)
}
