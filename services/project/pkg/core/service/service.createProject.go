package service

import "projects/pkg/core/models"

func (s *projectService) CreateProject(project models.Project, userId interface{}) (interface{}, error) {
	//TODO: logic for creating project and apllay validations
	if s.projectRepo.IsProjectExist(project, userId) {
		return nil, ErrProjectExists
	}

	projectId, err := s.projectRepo.CreateProject(project, userId)
	if err != nil {
		return nil, err
	}
	return projectId, nil

}
