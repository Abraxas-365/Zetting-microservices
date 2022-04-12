package service

import "projects/pkg/core/models"

func (s *projectService) GetProjectByProjectId(projectId interface{}) (*models.Project, error) {
	project, err := s.projectRepo.GetProjectByProjectId(projectId)
	if err != nil {
		return nil, err
	}
	return project, nil

}
