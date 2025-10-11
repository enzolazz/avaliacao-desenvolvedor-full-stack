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

func (s *AuthService) Login(username, password string) (string, string, *models.User, error) {
	user, _ := s.UserRepo.FindByUsername(username)
	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", "", nil, errors.New("nome de usuário ou senha inválida")
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID.Hex(),
		"username": user.Username,
		"exp":      time.Now().Add(config.Consts.AccessTokenExp).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID.Hex(),
		"username": user.Username,
		"exp":      time.Now().Add(config.Consts.RefreshTokenExp).Unix(),
	})

	accessSignedToken, err := accessToken.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", "", nil, err
	}

	refreshSignedToken, err := refreshToken.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", "", nil, err
	}

	return accessSignedToken, refreshSignedToken, user, nil
}

func (s *AuthService) RotateRefreshToken(refreshToken string) (string, string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenMalformed
		}
		return []byte(s.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return "", "", errors.New("invalid or expired refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("invalid token claims")
	}

	userIDRaw, ok := claims["user_id"]
	if !ok {
		return "", "", errors.New("user_id missing in claims")
	}

	usernameRaw, ok := claims["username"]
	if !ok {
		return "", "", errors.New("username missing in claims")
	}

	userID, ok := userIDRaw.(string)
	if !ok {
		return "", "", errors.New("user_id claim is not a string")
	}
	username, ok := usernameRaw.(string)
	if !ok {
		return "", "", errors.New("username claim is not a string")
	}

	newAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(config.Consts.AccessTokenExp).Unix(),
	})

	newRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(config.Consts.RefreshTokenExp).Unix(),
	})

	accessSigned, err := newAccess.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", "", err
	}
	refreshSigned, err := newRefresh.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", "", err
	}

	return accessSigned, refreshSigned, nil
}
