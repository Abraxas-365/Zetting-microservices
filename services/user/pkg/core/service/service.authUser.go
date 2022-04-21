package service

import (
	"fmt"
	"user/pkg/core/models"
)

func (r *userService) Auth(email models.Email, password models.Password) (models.User, error) {

	fmt.Println(password)
	user, err := r.userRepo.GetUserByEmail(email)
	if err != nil {
		return models.User{}, ErrUserNotFound
	}
	fmt.Println("user", user.Password)
	if !password.EqualTo(user.Password) {
		return models.User{}, ErrUnauthorized
	}

	return user, nil

}
