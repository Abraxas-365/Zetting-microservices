package ports

import "user/pkg/core/models"

type UserService interface {
	UpdateUser(userDataToUpdate *models.User, userId interface{}) error
	AddProjectToUser(projectData models.AddProjectToUser, document string) error
}
