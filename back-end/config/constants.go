package config

import (
	"log"
	"sync"
	"time"
)

type Constants struct {
	AccessTokenExp      time.Duration
	RefreshTokenExp     time.Duration
	HealthCheckInterval time.Duration
	IsAliveTimeout      time.Duration
	CorsMaxAge          time.Duration
	MaxInactiveFailures int
	MaxGoRoutines       int
}

var (
	Consts    *Constants
	constOnce sync.Once
)

func InitConstants() {
	constOnce.Do(func() {
		Consts = &Constants{
			AccessTokenExp:      1 * time.Hour,
			RefreshTokenExp:     30 * 24 * time.Hour,
			HealthCheckInterval: 10 * time.Minute,
			IsAliveTimeout:      5 * time.Second,
			CorsMaxAge:          12 * time.Hour,
			MaxInactiveFailures: 5,
			MaxGoRoutines:       10,
		}
		log.Println("Loaded constants successfully!")
	})
}
