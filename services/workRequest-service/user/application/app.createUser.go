package application

import "work-request/user/core/models"

func (a *userApplication) CreateUser(user models.User) error {
	//TODO ensure what happend when user already exist and mq reenquee the error
	if err := a.repo.CreateUser(user); err != nil {
		return err
	}
	return nil

}
