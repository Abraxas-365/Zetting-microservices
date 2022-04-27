package application

import (
	"errors"
	"user/pkg/core/models"
	"user/pkg/core/ports"
	"user/pkg/core/service"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
	ErrEmptyParams  = errors.New("Empty parameters")
	ErrUserExists   = errors.New("User already exists")
)

type UserApplication interface {
	UpdateUser(userDataToUpdate models.User, userId interface{}) error
	CreateUser(user models.User) (models.User, string, error)
	LoginUser(email models.Email, password models.Password) (models.User, string, error)
	IsUserExsist(email models.Email) error
	GetUserById(userId interface{}) (models.UserPublic, error)
}

type userApplication struct {
	userRepo    ports.UserRepository
	userService service.UserService
}

func NewUserApplication(userRepo ports.UserRepository, userService service.UserService) UserApplication {
	return &userApplication{
		userRepo,
		userService,
	}
}
