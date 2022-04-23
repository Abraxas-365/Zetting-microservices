package service

import (
	"errors"
	"user/pkg/core/models"
	"user/pkg/core/ports"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
	ErrEmptyParams  = errors.New("Empty parameters")
	ErrUserExists   = errors.New("User already exists")
)

type UserService interface {
	Auth(email models.Email, password models.Password) (models.User, error)
	CanCreateUser(user models.User) error
}

type userService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) UserService {
	return &userService{
		userRepo,
	}
}
