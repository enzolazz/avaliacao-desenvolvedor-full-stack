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
			FrontendServer: getEnv("FRONTEND_SERVER"),
			MongoURI:       getEnv("MONGODB_URI"),
			DBName:         getEnv("DB_NAME"),
			RedisURI:       getEnv("REDIS_URI"),
			ServerPort:     getEnv("SERVER_PORT"),
			JWTSecret:      getEnv("JWT_SECRET"),
		}

		log.Println("Loaded environment config successfully!")
	})
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}
