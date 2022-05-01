package service

import (
	"projects/project/core/models"
)

func (s *projectService) CanCreateProject(project models.Project) error {
	//TODO: logic for creating project and apllay validations
	if s.projectRepo.IsProjectExist(project) {
		return ErrProjectExists
	}
	return nil

}
