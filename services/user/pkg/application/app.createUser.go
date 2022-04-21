package application

import (
	"user/pkg/core/models"
)

func (r *userApplication) CreateUser(user models.User) (models.User, string, error) {

	if err := r.userService.CreateUser(user); err != nil {
		return models.User{}, "", err
	}
	return r.LoginUser(user.Contact.Email, user.Password)

}
