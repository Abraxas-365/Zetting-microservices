package application

import "projects/pkg/core/models"

func (s *projectApplication) GetProjectByProjectId(projectId interface{}) (models.Project, error) {
	project, err := s.projectRepo.GetProjectByProjectId(projectId)
	if err != nil {
		return models.Project{}, err
	}
	return project, nil

}
