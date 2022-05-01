package service

import (
	"errors"
	"projects/project/core/models"
	"projects/project/core/ports"
)

var (
	ErrUnauthorized  = errors.New("Unauthorized")
	ErrEmptyParams   = errors.New("Empty parameters")
	ErrProjectExists = errors.New("Project exists")
)

type ProjectService interface {
	CanCreateProject(project models.Project) error
}

type projectService struct {
	projectRepo ports.ProjectRepository
}

func NewProjectService(projectRepo ports.ProjectRepository) ProjectService {
	return &projectService{
		projectRepo,
	}
}
