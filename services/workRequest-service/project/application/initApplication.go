package application

import (
	"errors"
	"work-request/project/core/models"
	"work-request/project/core/ports"
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
