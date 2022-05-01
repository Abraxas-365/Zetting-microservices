package application

import "projects/workrequest/core/models"

func (a *workRequestApplication) WorkRequestAccepted(workrequest models.WorkRequest) error {

	if err := a.projectApp.AddUserToProject(workrequest.Worker, workrequest.Project); err != nil {
		return err
	}
	return nil

}
