package application

import (
	"projects/project/core/models"

	"github.com/google/uuid"
)

func (s *projectApplication) GetProjects(userId uuid.UUID, document string, page int) (models.Projects, error) {
	projects, err := s.projectRepo.GetProjects(userId, document, page)

	if err != nil || projects == nil {
		return []*models.Project{}, nil
	}
	return projects, nil

}
