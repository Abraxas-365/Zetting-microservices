package application

import "actor-service/user/core/models"

func (a *userApplication) FilterUsers(filter models.Filter, page int) (models.Users, error) {
	return a.repo.FilterUsers(filter, page)
}
