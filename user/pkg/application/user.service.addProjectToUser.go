package service

import "user/pkg/core/models"

func (s *userService) AddProjectToUser(projectData models.AddProjectToUser, document string) error {
	if err := s.userRepo.AddProjectToUser(projectData, document); err != nil {
		return err
	}
	return nil
}
