package ports

import (
	"work-request/user/core/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	UpdateUser(query interface{}, userId uuid.UUID) error
	CreateUser(user models.User) error
}
