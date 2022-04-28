package ports

import (
	"projects/pkg/core/events"
	"projects/pkg/core/models"

	"github.com/google/uuid"
)

type ProjectRepository interface {
	CreateProject(project models.Project) (events.Event, error)
	GetProjects(userId uuid.UUID, document string, page int) (models.Projects, error)
	GetProjectByProjectId(projectId uuid.UUID) (models.Project, error)
	AddUserToProject(userID uuid.UUID, projectId uuid.UUID, document string) (events.Event, error)
	IsProjectExist(newProject models.Project) bool
}
