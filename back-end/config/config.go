package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	MongoURI      string
	DBName        string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	ServerPort    string
	JWTSecret     string
}

func GetConfig() *Config {
	redisDB := 0
	if val := os.Getenv("REDIS_DB"); val != "" {
		db, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("Invalid REDIS_DB value: %v", err)
		}
		redisDB = db
	}

	cfg := &Config{
		MongoURI:      getEnv("MONGODB_URI"),
		DBName:        getEnv("DB_NAME"),
		RedisAddr:     getEnv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       redisDB,
		ServerPort:    getEnv("SERVER_PORT"),
		JWTSecret:     getEnv("JWT_SECRET"),
	}

	log.Println("Loaded environment config!")

	return cfg
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}
