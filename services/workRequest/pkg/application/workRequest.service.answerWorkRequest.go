package service

import "work-request/pkg/core/models"

func (s *workRequestService) AnswerWorkRequest(workRequest models.WorkRequest) error {
	/*answer workRequest*/

	if err := s.projectRepo.AnswerWorkRequest(workRequest); err != nil {
		return err
	}
	return nil

	/*TODO: function to call rabbitMQ and comunicate with project service(AddUserToProject)*/

}
