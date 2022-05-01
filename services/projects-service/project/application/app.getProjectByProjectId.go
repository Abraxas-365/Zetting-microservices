package application

import (
	"projects/project/core/models"

	"github.com/google/uuid"
)

func (s *projectApplication) GetProjectByProjectId(projectId uuid.UUID) (models.Project, error) {
	project, err := s.projectRepo.GetProjectByProjectId(projectId)
	if err != nil {
		return models.Project{}, err
	}
	return project, nil

}
