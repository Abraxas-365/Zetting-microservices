package application

import (
	"projects/project/core/models"

	"github.com/google/uuid"
)

func (s *projectApplication) GetProjectByProjectId(projectId uuid.UUID) (models.LookupProject, error) {
	project, err := s.projectRepo.GetProjectByProjectId(projectId)
	if err != nil {
		return models.LookupProject{}, err
	}
	return project, nil

}
