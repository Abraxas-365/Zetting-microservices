package application

import (
	"projects/project/core/models"
)

func (s *projectApplication) CreateProject(newProject models.Project) (models.Project, error) {

	if err := s.projectService.CanCreateProject(newProject); err != nil {
		return models.Project{}, err
	}

	event, err := s.projectRepo.CreateProject(newProject)
	if err != nil {
		return models.Project{}, err
	}
	s.projectMQPublisher.PublishEvent(event)
	s.userApp.AddProjectCount(newProject.Owner)

	return newProject, nil

}
