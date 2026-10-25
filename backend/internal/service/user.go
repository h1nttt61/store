package service

import (
	"errors"
	"kitten-backend/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	FindUserByUserName(username string) (*models.User, error)
	CreateUser(user *models.User) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(username, password string) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Password: string(hash),
	}
	return s.repo.CreateUser(user)
}

func (s *UserService) Login(username, password string) (*models.User, error) {
	user, err := s.repo.FindUserByUserName(username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("wrong password")
	}

	return user, nil
}
