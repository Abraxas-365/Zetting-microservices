package application

import (
	"work-request/pkg/core/models"
)

func (r *workRequestApplication) CreateWorkRequest(newWorkRequest models.WorkRequest) error {
	if err := r.projectService.CanCreateWorkRequest(newWorkRequest); err != nil {
		return err
	}

	event, err := r.projectRepo.CreateWorkRequest(newWorkRequest)
	if err != nil {
		return err
	}
	r.mqpublisher.PublishEvent(event)
	return nil

}
