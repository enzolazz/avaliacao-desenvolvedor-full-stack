package services

import (
	"errors"
	"time"

	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/models"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/repositories"
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

	// Always store the ObjectID as a hex string in the token
	token := newToken(user.ID.Hex(), user.Username)

	signedToken, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", nil, err
	}

	return signedToken, user, nil
}

func (s *AuthService) Refresh(oldToken string) (string, error) {
	token, err := jwt.Parse(oldToken, func(token *jwt.Token) (any, error) {
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
	userID, ok := userIDRaw.(string)
	if !ok {
		return "", errors.New("user_id claim is not a string")
	}

	usernameRaw, ok := claims["username"]
	if !ok {
		return "", errors.New("username claim missing")
	}
	username, ok := usernameRaw.(string)
	if !ok {
		return "", errors.New("username claim is not a string")
	}

	newToken := newToken(userID, username)
	signedToken, err := newToken.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func newToken(userID, username string) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
}
