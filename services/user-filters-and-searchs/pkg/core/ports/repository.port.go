package ports

import "user/pkg/core/models"

type UserRepository interface {
	GetUsersByProfession(profession string, page int) (models.Users, error)
	GetUsersByProject(projectId interface{}, document string) (models.Users, error)
}
