package application

import "github.com/google/uuid"

func (s *projectApplication) AddUserToProject(userID uuid.UUID, projectId uuid.UUID, document string) error {

	if _, err := s.projectRepo.AddUserToProject(userID, projectId, "workers"); err != nil {
		return err
	}

	return nil

}
