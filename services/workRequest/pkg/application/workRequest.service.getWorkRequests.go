package service

import "work-request/pkg/core/models"

func (r *workRequestService) GetWorkRequests(referenceId interface{}, status string, page int, number int, document string) (models.WorkRequests, error) {

	workRequests, err := r.projectRepo.GetWorkRequests(referenceId, status, page, number, document)
	if err != nil {
		return nil, err
	}
	return workRequests, nil

}
