package application

import "projects/pkg/core/models"

func (s *projectApplication) AddUserToProject(addUserData models.AddUserToProject, document string) error {

	if err := s.projectRepo.AddUserToProject(addUserData, document); err != nil {

		return err
	}

	return nil

}
