package application

import (
	"errors"
	"projects/user/core/models"
	"projects/user/core/ports"

	"github.com/google/uuid"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
)

type UserApplication interface {
	CreateUser(models.User) error
	UpdateUser(models.User) error
	AddProjectCount(userId uuid.UUID) error
	GetUserById(userId uuid.UUID) (models.User, error)
}
type userApplication struct {
	repo ports.UserRepository
}

func NewUserApplication(repo ports.UserRepository) UserApplication {
	return &userApplication{
		repo,
	}

}
