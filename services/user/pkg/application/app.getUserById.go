package application

import (
	"user/pkg/core/models"

	"github.com/google/uuid"
)

func (a userApplication) GetUserById(userId uuid.UUID) (models.UserPublic, error) {
	user, err := a.userRepo.GetUserById(userId)
	if err != nil {
		return models.UserPublic{}, err
	}
	return user, nil
}
