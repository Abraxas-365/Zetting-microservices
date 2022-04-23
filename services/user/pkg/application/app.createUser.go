package application

import (
	"user/pkg/core/models"
)

func (a *userApplication) CreateUser(user models.User) (models.User, string, error) {

	if err := a.userService.CanCreateUser(user); err != nil {
		return models.User{}, "", err
	}
	if err := a.userRepo.CreateUser(user); err != nil {
		return models.User{}, "", err
	}

	return a.LoginUser(user.Contact.Email, user.Password)

}
