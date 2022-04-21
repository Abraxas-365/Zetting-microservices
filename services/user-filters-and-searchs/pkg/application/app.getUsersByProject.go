package application

import "user/pkg/core/models"

func (s *userApplication) GetUsersByProject(projectId interface{}, document string) (models.Users, error) {

	users, err := s.userRepo.GetUsersByProject(projectId, document)
	if err != nil {
		return nil, err
	}
	return users, nil
}
