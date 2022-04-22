package service

import "work-request/pkg/core/models"

func (s *workRequestService) CreateWorkRequest(workRequest models.WorkRequest) error {
	if s.workRequestRepo.IsUserExsist(workRequest) {
		return ErrWorkRequestExists
	}
	if err := s.workRequestRepo.CreateWorkRequest(workRequest); err != nil {
		return err
	}
	return nil
}
