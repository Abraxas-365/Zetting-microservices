package ports

import "user/pkg/core/models"

type UserService interface {
	GetUsersByProfession(profession string, page int) (models.Users, error)
	GetUsersByProject(projectId interface{}, document string) (models.Users, error)
}
