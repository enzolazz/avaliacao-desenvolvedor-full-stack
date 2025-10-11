package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	FrontendServer string
	MongoURI       string
	DBName         string
	RedisAddr      string
	RedisPassword  string
	RedisDB        int
	ServerPort     string
	JWTSecret      string
}

type Constants struct {
	CookieExp           time.Duration
	HealthCheckInterval time.Duration
	IsAliveTimeout      time.Duration
	CorsMaxAge          time.Duration
	MaxInactiveFailures int
	MaxGoRoutines       int
}

func GetConfig() *Config {
	_ = godotenv.Load()

	redisDB := 0
	if val := os.Getenv("REDIS_DB"); val != "" {
		db, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("Invalid REDIS_DB value: %v", err)
		}
		redisDB = db
	}

	cfg := &Config{
		FrontendServer: getEnv("FRONTEND_SERVER"),
		MongoURI:       getEnv("MONGODB_URI"),
		DBName:         getEnv("DB_NAME"),
		RedisAddr:      getEnv("REDIS_ADDR"),
		RedisPassword:  os.Getenv("REDIS_PASSWORD"),
		RedisDB:        redisDB,
		ServerPort:     getEnv("SERVER_PORT"),
		JWTSecret:      getEnv("JWT_SECRET"),
	}

	log.Println("Loaded environment config!")

	return cfg
}

func GetConstants() *Constants {
	constants := &Constants{
		CookieExp:           1 * time.Hour,
		HealthCheckInterval: 10 * time.Minute,
		IsAliveTimeout:      5 * time.Second,
		CorsMaxAge:          12 * time.Hour,
		MaxInactiveFailures: 5,
		MaxGoRoutines:       10,
	}

	return constants
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}
