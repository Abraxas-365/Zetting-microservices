package ports

import (
	"user/user/core/events"
	"user/user/core/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	UpdateUser(query interface{}, userId uuid.UUID) error
	CreateUser(user models.User) (events.Event, error)
	GetUserByEmail(email models.Email) (models.User, error)
	IsUserExsist(email models.Email) bool
	GetUserById(userId uuid.UUID) (models.User, error)
}
