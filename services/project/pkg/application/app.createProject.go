package application

import (
	"fmt"
	"os"
	"projects/pkg/core/models"
)

func (s *projectApplication) CreateProject(newProject models.Project, userId interface{}) (models.Project, error) {

	if err := s.projectService.CanCreateProject(newProject, userId); err != nil {
		return models.Project{}, err
	}

	projectId, err := s.projectRepo.CreateProject(newProject, userId)
	if err != nil {
		return models.Project{}, err
	}
	if err := s.projectMQPublisher.NewProject(newProject, "Project", "newProject"); err != nil {
		// TODO:do something
		fmt.Println("Rabbit consumer closed - critical Error")
		os.Exit(1)
	}
	newProject.ID = projectId

	return newProject, nil

}
