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
	metricsRepo := repositories.NewMetricsRepository(db.Database.Collection("metrics"))

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo, jwtSecret)
	shortLinkService := services.NewShortLinkService(shortLinkRepo)
	metricsService := services.NewMetricsService(metricsRepo)

	userController := controllers.NewUserController(userService)
	profileController := controllers.NewProfileController(userService)
	authController := controllers.NewAuthController(authService)
	shortLinkController := controllers.NewShortLinkController(shortLinkService)
	metricsController := controllers.NewMetricsController(metricsService)
	redirectController := controllers.NewRedirectController(shortLinkService, metricsService)

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		auth.POST("/login", authController.Login)
		auth.POST("/register", userController.Register)

		profile := api.Group("/profile")
		profile.Use(middleware.JWTMiddleware(jwtSecret))
		{
			profile.GET("/me", profileController.Me)
		}

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
