package ports

import (
	"work-request/project/core/models"
)

type ProjectRepository interface {
	CreateProject(project models.Project) error
}
