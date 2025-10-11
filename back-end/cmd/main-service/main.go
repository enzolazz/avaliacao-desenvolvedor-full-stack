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
)

func main() {
	db.ConnectMongoDB(config.Cfg.MongoURI, config.Cfg.DBName)
	linksCollection := db.Client.Database(config.Cfg.DBName).Collection("shortlinks")

	ps := pubsub.NewRedisPubSub()

	unsubscribe, err := handlers.HandleLinkStatusUpdates(ps, linksCollection)
	if err != nil {
		log.Fatal("Failed to start healthcheck handler:", err)
	}
	defer unsubscribe()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.Cfg.FrontendServer},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           config.Consts.CorsMaxAge,
	}))

	routes.RegisterRoutes(r, config.Cfg.JWTSecret)
	r.GET("/updates/ws", middleware.JWTMiddleware(config.Cfg.JWTSecret), handlers.WebSocketHandler(ps))

	r.Run(":" + config.Cfg.ServerPort)
}
