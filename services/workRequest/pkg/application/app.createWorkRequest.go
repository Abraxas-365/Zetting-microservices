package application

import (
	"fmt"
	"os"
	"work-request/pkg/core/models"
)

func (r *workRequestApplication) CreateWorkRequest(newWorkRequest models.WorkRequest) (models.WorkRequest, error) {
	newWorkRequest.Status = "P"
	WorkRequest, err := r.projectRepo.CreateWorkRequest(newWorkRequest)
	if err != nil {
		return models.WorkRequest{}, err
	}
	if err := r.mqpublisher.NewWorkRequest(WorkRequest, "WorkRequest", "new_workrequest"); err != nil {
		// TODO:do something
		fmt.Println("Rabbit consumer closed - critical Error")
		os.Exit(1)
	}

	return WorkRequest, nil

}
