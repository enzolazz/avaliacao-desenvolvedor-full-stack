package services

import (
	"errors"
	"time"

	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/models"
	"github.com/enzolazz/avaliacao-desenvolvedor-full-stack/back-end/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user models.User) error {
	existing, _ := s.Repo.FindByUsername(user.Username)
	if existing != nil {
		return errors.New("nome de usuário já existe")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	return s.Repo.Create(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.FindAll()
}

func (s *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	return s.Repo.FindByID(id)
}

func (s *UserService) LoginHandler(username, password, jwtSecret string) (string, error) {
	user, _ := s.Repo.FindByUsername(username)
	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("nome de usuário ou senha inválida")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID.String(),
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(jwtSecret))
}
