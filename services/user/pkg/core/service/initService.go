package service

import (
	"errors"
	"user/pkg/core/models"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
	ErrEmptyParams  = errors.New("Empty parameters")
	ErrUserExists   = errors.New("User already exists")
)

type UserService interface {
	Auth(email models.Email, password models.Password) (models.User, error)
	CreateUser(user models.User) error
}

type ServiceRepository interface {
	GetUserByEmail(email models.Email) (models.User, error)
	CreateUser(user models.User) error
	IsUserExsist(user models.User) bool
}
type userService struct {
	userRepo ServiceRepository
}

func NewUsertService(userRepo ServiceRepository) UserService {
	return &userService{
		userRepo,
	}
}
