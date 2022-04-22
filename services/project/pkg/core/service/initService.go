package service

import (
	"errors"
	"projects/pkg/core/models"
)

var (
	ErrUnauthorized  = errors.New("Unauthorized")
	ErrEmptyParams   = errors.New("Empty parameters")
	ErrProjectExists = errors.New("Project exists")
)

type ProjectService interface {
	CreateProject(project models.Project, userId interface{}) (interface{}, error)
}

type ServiceRepository interface {
	CreateProject(project models.Project, userId interface{}) (interface{}, error)
	IsProjectExist(newProject models.Project, userId interface{}) bool
}
type projectService struct {
	projectRepo ServiceRepository
}

func NewProjectService(projectRepo ServiceRepository) ProjectService {
	return &projectService{
		projectRepo,
	}
}
