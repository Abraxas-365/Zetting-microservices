package application

import (
	"github.com/google/uuid"
)

func (s *projectApplication) AddUserToProject(user uuid.UUID, project uuid.UUID) error {

	if _, err := s.projectRepo.AddUserToProject(user, project); err != nil {
		return err
	}

	return nil

}
