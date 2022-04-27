package application

import "user/pkg/core/models"

func (a userApplication) GetUserById(userId interface{}) (models.UserPublic, error) {
	user, err := a.userRepo.GetUserById(userId)
	if err != nil {
		return models.UserPublic{}, err
	}
	return user, nil
}
