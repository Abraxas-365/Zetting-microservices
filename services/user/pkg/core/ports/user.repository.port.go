package ports

import "user/pkg/core/models"

type UserRepository interface {
	CreateUser(user models.User) (*models.User, error)
	UpdateUser(query interface{}, userId interface{}) error
	CheckEmailExist(email string) (*models.User, error)
	GetUserById(userId interface{}) (*models.User, error)
	GetUsersByProfession(profession string, page int) (models.Users, error)
	GetUsersByProject(projectId interface{}, document string) (models.Users, error)
	AddProjectToUser(projectData models.AddProjectToUser, document string) error
}
