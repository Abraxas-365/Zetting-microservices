package ports

import "projects/pkg/core/models"

type ProjectRepository interface {
	CreateProject(project models.Project, userId interface{}) (interface{}, error)
	GetProjects(userId interface{}, document string, page int) (models.Projects, error)
	GetProjectByProjectId(projectId interface{}) (models.Project, error)
	AddUserToProject(addUserData models.AddUserToProject, document string) error
	IsProjectExist(newProject models.Project, userId interface{}) bool
}
