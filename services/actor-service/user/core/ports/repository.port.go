package ports

import (
	"actor-service/user/core/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	UpdateUser(query interface{}, userId uuid.UUID) error
	CreateUser(user models.User) error
	GetUsers(page int) (models.Users, error)
	GetUserById(user uuid.UUID) (models.User, error)
	FilterUsers(filter models.Filter, page int) (models.Users, error)
}
