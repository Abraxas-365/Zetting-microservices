package application

import (
	"projects/user/core/models"

	"github.com/google/uuid"
)

func (a *userApplication) GetUserById(userId uuid.UUID) (models.User, error) {
	return a.repo.GetUserById(userId)
}
