package application

import (
	"errors"
	"notifications/user/core/models"
	"notifications/user/core/ports"
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
