package application

import (
	"errors"
	"projects/pkg/core/models"
	"projects/pkg/core/ports"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
)

type ProjectApplication interface {
	CreateProject(newProject *models.Project, userId interface{}) (interface{}, error)
	GetProjects(userId interface{}, document string, page int) (models.Projects, error)
	AddUserToProject(addUserData models.AddUserToProject, document string) error
	GetProjectByProjectId(projectId interface{}) (models.Project, error)
}
type projectApplication struct {
	projectRepo        ports.ProjectRepository
	projectMQPublisher ports.ProjectMQPublisher
}

func NewProjectApplication(projectRepo ports.ProjectRepository, mqPublisher ports.ProjectMQPublisher) ProjectApplication {
	return &projectApplication{
		projectRepo,
		mqPublisher,
	}

}
