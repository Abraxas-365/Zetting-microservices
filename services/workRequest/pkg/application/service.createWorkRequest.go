package service

import (
	"work-request/pkg/core/models"
)

func (r *workRequestService) CreateWorkRequest(newWorkRequest models.WorkRequest) (*models.WorkRequest, error) {
	newWorkRequest.Status = "P"
	WorkRequest, err := r.projectRepo.CreateWorkRequest(newWorkRequest)
	if err != nil {
		return nil, err
	}

	if err := r.mqpublisher.Publish(*WorkRequest); err != nil {
		//TODO: si falla eliminar el work request de la base de datos
		return WorkRequest, err
	}

	return WorkRequest, nil

}
