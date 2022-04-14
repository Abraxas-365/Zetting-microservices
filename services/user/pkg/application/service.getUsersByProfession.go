package service

import "user/pkg/core/models"

func (r *userService) GetUsersByProfession(profession string, page int) (models.Users, error) {
	users, err := r.userRepo.GetUsersByProfession(profession, page)
	if err != nil {
		return nil, err
	}
	return users, nil
}
