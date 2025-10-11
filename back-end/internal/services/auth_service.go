package services

import (
	"errors"
	"time"
	"url-shortener/back-end/config"
	"url-shortener/back-end/internal/models"
	"url-shortener/back-end/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo  *repositories.UserRepository
	JWTSecret string
}

func NewAuthService(repo *repositories.UserRepository, secret string) *AuthService {
	return &AuthService{UserRepo: repo, JWTSecret: secret}
}

func (s *AuthService) Login(username, password string) (string, *models.User, error) {
	user, _ := s.UserRepo.FindByUsername(username)
	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", nil, errors.New("nome de usuário ou senha inválida")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID.Hex(),
		"username": user.Username,
		"exp":      time.Now().Add(config.GetConstants().CookieExp).Unix(),
	})

	signedToken, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", nil, err
	}

	return signedToken, user, nil
}

func (s *AuthService) Refresh(oldToken string) (string, error) {
	token, err := jwt.Parse(oldToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenMalformed
		}
		return []byte(s.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userIDRaw, ok := claims["user_id"]
	if !ok {
		return "", errors.New("user_id claim missing")
	}

	usernameRaw, ok := claims["username"]
	if !ok {
		return "", errors.New("username claim missing")
	}

	username, ok := usernameRaw.(string)
	if !ok {
		return "", errors.New("username claim is not a string")
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userIDRaw,
		"username": username,
		"exp":      time.Now().Add(config.GetConstants().CookieExp).Unix(),
	})

	signedToken, err := newToken.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
