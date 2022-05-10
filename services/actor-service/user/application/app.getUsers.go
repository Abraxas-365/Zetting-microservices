package application

import "actor-service/user/core/models"

func (a *userApplication) GetUsers() (models.Users, error) {
	return a.repo.GetUsers()
}
