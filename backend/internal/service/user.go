package service

import (
	"errors"
	"kitten-backend/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	FindUserByUserName(username string) (*models.User, error)
	CreateUser(user *models.User) error
	GetUserByID(userID uint) (*models.User, error)
	UpdateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(username, email, password string) error {
	if _, err := s.repo.FindUserByUserName(username); err == nil {
		return errors.New("username already exists")
	}
	if _, err := s.repo.FindUserByEmail(email); err == nil {
		return errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hash),
		Role:     "user",
	}
	return s.repo.CreateUser(user)
}

func (s *UserService) Login(username, password string) (*models.User, error) {
	user, err := s.repo.FindUserByUserName(username)
	if err != nil {
		user, err = s.repo.FindUserByEmail(username)
		if err != nil {
			return nil, errors.New("user not found")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("wrong password")
	}
	return user, nil
}

func (s *UserService) MakeAdmin(userID uint) error {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return err
	}
	user.Role = "admin"
	return s.repo.UpdateUser(user)
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	return s.repo.FindUserByUserName(username)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.FindUserByEmail(email)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}
