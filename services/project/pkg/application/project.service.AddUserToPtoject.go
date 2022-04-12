package service

import "projects/pkg/core/models"

func (s *projectService) AddUserToProject(addUserData models.AddUserToProject, document string) error {

	if err := s.projectRepo.AddUserToProject(addUserData, document); err != nil {

		return err
	}

	return nil

}
