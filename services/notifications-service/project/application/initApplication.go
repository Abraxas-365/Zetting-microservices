package application

import (
	"errors"
	"notifications/project/core/models"
	"notifications/project/core/ports"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
)

type ProjectApplication interface {
	CreateProject(models.Project) error
}
type projectApplication struct {
	repo ports.ProjectRepository
}

func NewProjectApplication(repo ports.ProjectRepository) ProjectApplication {
	return &projectApplication{
		repo,
	}

}
