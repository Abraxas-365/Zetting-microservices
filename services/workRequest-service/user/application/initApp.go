package application

import (
	"errors"
	"work-request/user/core/models"
	"work-request/user/core/ports"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
)

type UserApplication interface {
	CreateUser(models.User) error
	UpdateUser(models.User) error
}
type userApplication struct {
	repo ports.UserRepository
}

func NewUserApplication(repo ports.UserRepository) UserApplication {
	return &userApplication{
		repo,
	}

}
