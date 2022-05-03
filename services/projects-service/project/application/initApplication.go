package application

import (
	"errors"
	"projects/project/core/models"
	"projects/project/core/ports"
	"projects/project/core/service"

	"github.com/google/uuid"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
)

type ProjectApplication interface {
	CreateProject(newProject models.Project) (models.Project, error)
	GetProjects(userId uuid.UUID, document string, page int) (models.LookupProjects, error)
	AddUserToProject(user uuid.UUID, project uuid.UUID) error
	GetProjectByProjectId(projectId uuid.UUID) (models.LookupProject, error)
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
