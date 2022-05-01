package application

import (
	"work-request/workrequest/core/models"

	"github.com/google/uuid"
)

func (s *workRequestApplication) AnswerWorkRequest(workRequest uuid.UUID, workerId uuid.UUID, status models.Status) error {

	//VAlidate the data
	if err := status.Validate(); err != nil {
		return err
	}
	//Safe the data via repository
	event, err := s.projectRepo.AnswerWorkRequest(workRequest, workerId, string(status))
	if err != nil {
		return err
	}
	//sent to rabbitMQ
	s.mqpublisher.PublishEvent(event)

	return nil

}
