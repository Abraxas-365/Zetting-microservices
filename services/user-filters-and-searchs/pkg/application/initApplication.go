package application

import (
	"errors"
	"user/pkg/core/models"
	"user/pkg/core/ports"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
	ErrEmptyParams  = errors.New("Empty parameters")
)

type UserApplication interface {
	GetUsersByProfession(profession string, page int) (models.Users, error)
	GetUsersByProject(projectId interface{}, document string) (models.Users, error)
}
type userApplication struct {
	userRepo ports.UserRepository
}

func NewUserApplication(userRepo ports.UserRepository) UserApplication {
	return &userApplication{
		userRepo,
	}
}
