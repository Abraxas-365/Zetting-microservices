package service

import "work-request/pkg/core/models"

func (s *workRequestService) CreateWorkRequest(workRequest models.WorkRequest) error {
	if s.workRequestRepo.IsWorkrequestExist(workRequest) {
		return ErrWorkRequestExists
	}
	return nil
}
