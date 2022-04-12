package service

import "projects/pkg/core/models"

func (s *projectService) CreateProject(newProject *models.Project, userId interface{}) (interface{}, error) {

	projectId, err := s.projectRepo.CreateProject(newProject, userId)
	if err != nil {
		return "", err
	}

	return projectId, nil

}
