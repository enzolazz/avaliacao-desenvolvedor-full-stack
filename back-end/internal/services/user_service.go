package services

import (
	"errors"

	"url-shortener/back-end/internal/models"
	"url-shortener/back-end/internal/repositories"
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
