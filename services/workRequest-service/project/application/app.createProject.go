package application

import "work-request/project/core/models"

func (a *projectApplication) CreateProject(project models.Project) error {
	if err := a.repo.CreateProject(project); err != nil {
		return err
	}
	return nil

}
