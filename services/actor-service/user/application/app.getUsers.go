package application

import "actor-service/user/core/models"

func (a *userApplication) GetUsers(page int) (models.Users, error) {
	return a.repo.GetUsers(page)
}
