package utils

import (
	"strings"
	"url-shortener/back-end/config"
)

func GetDomain() (bool, string) {
	frontendURL := config.Cfg.FrontendServer

	secure := false
	domain := frontendURL

	if strings.HasPrefix(frontendURL, "https://") {
		secure = true
		domain = strings.TrimPrefix(frontendURL, "https://")
	} else if strings.HasPrefix(frontendURL, "http://") {
		secure = false
		domain = strings.TrimPrefix(frontendURL, "http://")
	}

	domain = strings.TrimSuffix(domain, "/")

	return secure, domain
}
