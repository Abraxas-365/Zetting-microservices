package application

import (
	"errors"
	"projects/pkg/core/models"
	"projects/pkg/core/ports"
	"projects/pkg/core/service"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
)

type ProjectApplication interface {
	CreateProject(newProject models.Project, userId interface{}) (models.Project, error)
	GetProjects(userId interface{}, document string, page int) (models.Projects, error)
	AddUserToProject(userID interface{}, projectId interface{}, document string) error
	GetProjectByProjectId(projectId interface{}) (models.Project, error)
}
type projectApplication struct {
	projectRepo        ports.ProjectRepository
	projectMQPublisher ports.ProjectMQPublisher
	projectService     service.ProjectService
}

func NewProjectApplication(projectRepo ports.ProjectRepository, mqPublisher ports.ProjectMQPublisher, projectService service.ProjectService) ProjectApplication {
	return &projectApplication{
		projectRepo,
		mqPublisher,
		projectService,
	}

}
