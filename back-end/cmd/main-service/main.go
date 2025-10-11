package main

import (
	"log"
	"url-shortener/back-end/config"
	"url-shortener/back-end/internal/db"
	"url-shortener/back-end/internal/handlers"
	"url-shortener/back-end/internal/middleware"
	"url-shortener/back-end/internal/pubsub"
	"url-shortener/back-end/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env from project root!")
	}

	cfg := config.GetConfig()

	db.ConnectMongoDB(cfg.MongoURI, cfg.DBName)
	linksCollection := db.Client.Database(cfg.DBName).Collection("shortlinks")

	ps := pubsub.NewRedisPubSub(cfg)

	unsubscribe, err := handlers.HandleLinkStatusUpdates(ps, linksCollection)
	if err != nil {
		log.Fatal("Failed to start healthcheck handler:", err)
	}
	defer unsubscribe()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.GetConfig().FrontendServer},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           config.GetConstants().CorsMaxAge,
	}))

	routes.RegisterRoutes(r, cfg.JWTSecret)
	r.GET("/updates/ws", middleware.JWTMiddleware(cfg.JWTSecret), handlers.WebSocketHandler(ps))

	r.Run(":" + cfg.ServerPort)
}
