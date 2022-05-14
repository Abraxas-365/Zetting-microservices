package ports

import (
	"projects/user/core/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	UpdateUser(query interface{}, userId uuid.UUID) error
	CreateUser(user models.User) error
	AddProjectCount(userId uuid.UUID) error
	GetUserById(userId uuid.UUID) (models.User, error)
}
