package service

import "work-request/pkg/core/models"

func (s *workRequestService) AnswerWorkRequest(workRequest models.WorkRequest) error {

	if err := s.projectRepo.AnswerWorkRequest(workRequest); err != nil {
		return err
	}
	if workRequest.Status == "A" {
		//sent to rabbitMQ add user to worker
		// eliminate the projectRepo part that do that
	}

	return nil

}
