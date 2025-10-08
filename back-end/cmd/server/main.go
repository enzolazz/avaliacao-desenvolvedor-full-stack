package main

import (
	"log"
	"time"

	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/config"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/db"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/routes"
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

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterAuthRoutes(r, cfg.JWTSecret)

	r.Run(":" + cfg.ServerPort)
}
