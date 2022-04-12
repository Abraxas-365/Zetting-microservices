package service

import (
	"errors"
	"projects/pkg/core/ports"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
)

type projectService struct {
	projectRepo ports.ProjectRepository
}

func NewProjectService(projectRepo ports.ProjectRepository) ports.ProjectService {
	return &projectService{
		projectRepo,
	}

}
