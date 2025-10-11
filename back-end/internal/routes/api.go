package routes

import (
	"url-shortener/back-end/internal/controllers"
	"url-shortener/back-end/internal/db"
	"url-shortener/back-end/internal/middleware"
	"url-shortener/back-end/internal/repositories"
	"url-shortener/back-end/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, jwtSecret string) {
	userRepo := repositories.NewUserRepository(db.Database.Collection("users"))
	shortLinkRepo := repositories.NewShortLinkRepository(db.Database.Collection("shortlinks"))
	metricsRepo := repositories.NewMetricsRepository(db.Database.Collection("metrics"))

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo, jwtSecret)
	shortLinkService := services.NewShortLinkService(shortLinkRepo)
	metricsService := services.NewMetricsService(metricsRepo)

	userController := controllers.NewUserController(userService, authService)
	authController := controllers.NewAuthController(authService)
	shortLinkController := controllers.NewShortLinkController(shortLinkService)
	metricsController := controllers.NewMetricsController(metricsService)
	redirectController := controllers.NewRedirectController(shortLinkService, metricsService)

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		auth.POST("/login", authController.Login)
		auth.POST("/refresh", authController.Refresh)
		auth.POST("/logout", authController.Logout)
		auth.POST("/register", userController.Register)

		users := api.Group("/users")
		users.Use(middleware.AllowOnlyLocalhost())
		{
			users.GET("", userController.GetAll)
		}

		shortlinks := api.Group("/shorten")
		shortlinks.Use(middleware.JWTMiddleware(jwtSecret))
		{
			shortlinks.POST("", shortLinkController.Create)
			shortlinks.GET("", shortLinkController.GetAllByUser)
		}

		metrics := api.Group("/metrics")
		metrics.Use(middleware.JWTMiddleware(jwtSecret))
		{
			metrics.GET("/last-hour/:id", metricsController.LastHour)
			metrics.GET("/last-day/:id", metricsController.LastDay)
			metrics.GET("/last-month/:id", metricsController.LastMonth)
		}

	}
	// Public search route
	r.GET("/redirect/:id", redirectController.HandleRedirect)
}
