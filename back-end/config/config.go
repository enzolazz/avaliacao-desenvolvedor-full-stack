package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	FrontendServer string
	MongoURI       string
	DBName         string
	RedisURI       string
	ServerPort     string
	JWTSecret      string
}

var (
	Cfg     *Config
	cfgOnce sync.Once
)

func InitConfig() {
	cfgOnce.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using system environment variables")
		}

		Cfg = &Config{
			FrontendServer: getRequiredEnv("FRONTEND_SERVER"),
			MongoURI:       getRequiredEnv("MONGODB_URI"),
			DBName:         getOptionalEnv("DB_NAME", "url_shortener"),
			RedisURI:       getRequiredEnv("REDIS_URI"),
			ServerPort:     getOptionalEnv("SERVER_PORT", "8080"),
			JWTSecret:      getRequiredEnv("JWT_SECRET"),
		}

		log.Println("Loaded environment config successfully!")
	})
}

func getRequiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}

func getOptionalEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
