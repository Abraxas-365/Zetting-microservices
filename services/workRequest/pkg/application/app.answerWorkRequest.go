package application

import (
	"work-request/pkg/core/models"
)

func (s *workRequestApplication) AnswerWorkRequest(workRequest models.WorkRequest) error {

	//VAlidate the data
	if err := workRequest.Validate(); err != nil {
		return err
	}
	//Safe the data via repository
	event, err := s.projectRepo.AnswerWorkRequest(workRequest)
	if err != nil {
		return err
	}
	//sent to rabbitMQ
	s.mqpublisher.AnswerWorkRequest(event)

	return nil

}
