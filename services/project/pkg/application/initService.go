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
	projectRepo        ports.ProjectRepository
	projectMQPublisher ports.ProjectMQPublisher
}

func NewProjectService(projectRepo ports.ProjectRepository, mqPublisher ports.ProjectMQPublisher) ports.ProjectService {
	return &projectService{
		projectRepo,
		mqPublisher,
	}

}
