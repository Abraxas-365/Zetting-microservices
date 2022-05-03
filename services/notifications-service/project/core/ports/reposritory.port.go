package ports

import (
	"notifications/project/core/models"
)

type ProjectRepository interface {
	CreateProject(project models.Project) error
}
