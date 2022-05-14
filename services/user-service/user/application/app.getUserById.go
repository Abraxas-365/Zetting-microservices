package application

import (
	"user/user/core/models"

	"github.com/google/uuid"
)

func (a *userApplication) GetUserById(userId uuid.UUID) (models.User, error) {
	user, err := a.userRepo.GetUserById(userId)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
