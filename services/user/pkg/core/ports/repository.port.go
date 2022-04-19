package ports

import "user/pkg/core/models"

type UserRepository interface {
	UpdateUser(query interface{}, userId interface{}) error
	AddProjectToUser(projectData models.AddProjectToUser, document string) error
}
