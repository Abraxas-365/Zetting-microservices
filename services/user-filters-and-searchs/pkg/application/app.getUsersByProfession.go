package application

import "user/pkg/core/models"

func (r *userApplication) GetUsersByProfession(profession string, page int) (models.Users, error) {
	users, err := r.userRepo.GetUsersByProfession(profession, page)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return models.Users{}, nil
	}
	return users, nil
}
