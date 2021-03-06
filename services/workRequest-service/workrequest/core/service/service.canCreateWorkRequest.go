package service

import "work-request/workrequest/core/models"

func (s *workRequestService) CanCreateWorkRequest(workRequest models.WorkRequest) error {
	if s.workRequestRepo.IsWorkrequestExist(workRequest) {
		return ErrWorkRequestExists
	}
	return nil
}
