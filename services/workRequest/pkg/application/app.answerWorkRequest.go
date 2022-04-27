package application

import (
	"work-request/pkg/core/events"
	"work-request/pkg/core/models"
)

func (s *workRequestApplication) AnswerWorkRequest(workRequest models.WorkRequest) error {

	//VAlidate the data
	if err := workRequest.Validate(); err != nil {
		return err
	}
	//Safe the data via repository
	workRequest, err := s.projectRepo.AnswerWorkRequest(workRequest)
	if err != nil {
		return err
	}
	//sent to rabbitMQ
	switch workRequest.Status {
	case "A":
		s.mqpublisher.AnswerWorkRequest(events.WorkrequestAccepted{
			ID:      workRequest.ID,
			Owner:   workRequest.Owner,
			Worker:  workRequest.Worker,
			Project: workRequest.Project,
		})
	case "B":
		s.mqpublisher.AnswerWorkRequest(events.WorkrequestDenied{
			ID:      workRequest.ID,
			Owner:   workRequest.Owner,
			Worker:  workRequest.Worker,
			Project: workRequest.Project,
		})
	}

	return nil

}
