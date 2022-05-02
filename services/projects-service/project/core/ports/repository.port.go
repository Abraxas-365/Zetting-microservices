package ports

import (
	"projects/project/core/events"
	"projects/project/core/models"

	"github.com/google/uuid"
)

type ProjectRepository interface {
	CreateProject(project models.Project) (events.Event, error)
	GetProjects(userId uuid.UUID, document string, page int) (models.LookupProjects, error)
	GetProjectByProjectId(projectId uuid.UUID) (models.LookupProject, error)
	AddUserToProject(user uuid.UUID, project uuid.UUID) (events.Event, error)
	AddWorkRequestToProject(projectId uuid.UUID, workRequestId uuid.UUID) (events.Event, error)
	IsProjectExist(newProject models.Project) bool
}
