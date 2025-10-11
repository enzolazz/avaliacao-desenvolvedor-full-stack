package utils

import (
	"net/http"
	"url-shortener/back-end/config"
)

func IsURLAlive(url string) bool {
	client := http.Client{
		Timeout: config.GetConstants().HealthCheckInterval,
	}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode >= 200 && resp.StatusCode < 400
}
