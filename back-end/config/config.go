package config

import (
	"log"
	"os"
)

type Config struct {
	MongoURI   string
	DBName     string
	RedisAddr  string
	ServerPort string
	JWTSecret  string
}

func GetConfig() *Config {
	config := &Config{
		MongoURI:   getEnv("MONGODB_URI"),
		DBName:     getEnv("DB_NAME"),
		RedisAddr:  getEnv("REDIS_ADDR"),
		ServerPort: getEnv("SERVER_PORT"),
		JWTSecret:  getEnv("JWT_SECRET"),
	}

	log.Println("Loaded environment config!")

	return config
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}

	return value
}
