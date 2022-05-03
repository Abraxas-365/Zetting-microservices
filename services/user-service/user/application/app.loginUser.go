package application

import (
	// "user/internal/auth"
	"user/internal/auth"
	"user/user/core/models"
)

func (r *userApplication) LoginUser(email models.Email, password models.Password) (models.User, string, error) {

	user, err := r.userService.Auth(email, password)
	if err != nil {
		return models.User{}, "", err
	}
	token, err := auth.GereteToken(string(user.Contact.Email), user.ID)
	if err != nil {
		return models.User{}, "", err
	}

	return user, token, nil
}
