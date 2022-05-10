package application

import (
	"fmt"
	"user/user/core/models"
)

func (a *userApplication) CreateUser(user models.User) (models.User, string, error) {
	if err := a.userService.CanCreateUser(user); err != nil {
		return models.User{}, "", err
	}
	event, err := a.userRepo.CreateUser(user)
	if err != nil {
		return models.User{}, "", err
	}
	fmt.Println(event)
	a.userPublish.PublishEvent(event)

	return a.LoginUser(user.Contact.Email, user.Password)

}
