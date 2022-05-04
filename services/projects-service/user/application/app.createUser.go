package application

import "projects/user/core/models"

func (a *userApplication) CreateUser(user models.User) error {
	//TODO Check if user exist and if exist return nil
	if err := a.repo.CreateUser(user); err != nil {
		return err
	}
	return nil

}
