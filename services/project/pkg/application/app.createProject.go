package application

import (
	"fmt"
	"os"
	"projects/pkg/core/models"
)

func (s *projectApplication) CreateProject(newProject models.Project, userId interface{}) (interface{}, error) {

	projectId, err := s.projectService.CreateProject(newProject, userId)
	if err != nil {
		return "", err
	}
	if err := s.projectMQPublisher.NewProject(newProject, "Project", "newProject"); err != nil {
		// TODO:do something
		fmt.Println("Rabbit consumer closed - critical Error")
		os.Exit(1)
	}

	return projectId, nil

}
