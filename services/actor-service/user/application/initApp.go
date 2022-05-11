package application

import (
	"actor-service/user/core/models"
	"actor-service/user/core/ports"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
)

type UserApplication interface {
	CreateUser(user models.User) error
	UpdateUser(updated models.User, userId uuid.UUID) error
	GetUsers(page int) (models.Users, error)
	FilterUsers(filtro models.Filter, page int) (models.Users, error)
}
type userApplication struct {
	repo ports.UserRepository
}

func NewUserApplication(repo ports.UserRepository) UserApplication {
	return &userApplication{
		repo,
	}

}
