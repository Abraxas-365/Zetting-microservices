package ports

import "projects/pkg/core/models"

type ProjectService interface {
	CreateProject(newProject *models.Project, userId interface{}) (interface{}, error)
	GetProjects(userId interface{}, document string, page int) (models.Projects, error)
	AddUserToProject(addUserData models.AddUserToProject, document string) error
	GetProjectByProjectId(projectId interface{}) (*models.Project, error)
}
