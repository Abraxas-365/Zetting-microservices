package service

import (
	"fmt"
	"os"
	"work-request/pkg/core/models"
)

func (r *workRequestService) CreateWorkRequest(newWorkRequest models.WorkRequest) (*models.WorkRequest, error) {
	newWorkRequest.Status = "P"
	WorkRequest, err := r.projectRepo.CreateWorkRequest(newWorkRequest)
	if err != nil {
		return nil, err
	}
	if err := r.mqpublisher.PublishNewWorkRequest(*WorkRequest, "WorkRequest", "workrequest"); err != nil {
		// TODO:do something
		fmt.Println("Rabbit consumer closed - critical Error")
		os.Exit(1)
	}

	return WorkRequest, nil

}
