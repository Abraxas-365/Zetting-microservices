package application

import (
	"errors"
	"projects/pkg/core/models"
	"projects/pkg/core/ports"
	"projects/pkg/core/service"

	"github.com/google/uuid"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
)

type ProjectApplication interface {
	CreateProject(newProject models.Project) (models.Project, error)
	GetProjects(userId uuid.UUID, document string, page int) (models.Projects, error)
	AddUserToProject(userID uuid.UUID, projectId uuid.UUID, document string) error
	GetProjectByProjectId(projectId uuid.UUID) (models.Project, error)
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
