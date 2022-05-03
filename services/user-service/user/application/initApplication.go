package application

import (
	"errors"
	"user/user/core/models"
	"user/user/core/ports"
	"user/user/core/service"

	"github.com/google/uuid"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
	ErrEmptyParams  = errors.New("Empty parameters")
	ErrUserExists   = errors.New("User already exists")
)

type UserApplication interface {
	UpdateUser(userDataToUpdate models.User, userId uuid.UUID) error
	CreateUser(user models.User) (models.User, string, error)
	LoginUser(email models.Email, password models.Password) (models.User, string, error)
	IsUserExsist(email models.Email) error
	GetUserById(userId uuid.UUID) (models.UserPublic, error)
}

type userApplication struct {
	userRepo    ports.UserRepository
	userService service.UserService
	userPublish ports.UserMQPublisher
}

func NewUserApplication(userRepo ports.UserRepository, userService service.UserService, userPublish ports.UserMQPublisher) UserApplication {
	return &userApplication{
		userRepo,
		userService,
		userPublish,
	}
}
