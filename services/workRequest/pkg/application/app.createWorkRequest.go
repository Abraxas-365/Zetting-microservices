package application

import (
	"work-request/pkg/core/events"
	"work-request/pkg/core/models"
)

func (r *workRequestApplication) CreateWorkRequest(newWorkRequest models.WorkRequest) error {
	if err := r.projectService.CreateWorkRequest(newWorkRequest); err != nil {
		return err
	}

	workRequest, err := r.projectRepo.CreateWorkRequest(newWorkRequest)
	if err != nil {
		return err
	}
	r.mqpublisher.NewWorkRequest(events.WorkrequestCreated{
		ID:      workRequest.ID,
		Owner:   workRequest.Owner,
		Worker:  workRequest.Worker,
		Project: workRequest.Project,
	})
	return nil

}
