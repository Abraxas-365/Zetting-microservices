package application

import (
	"work-request/pkg/core/models"

	"github.com/google/uuid"
)

func (r *workRequestApplication) GetWorkRequests(referenceId uuid.UUID, status string, page int, number int, document string) (models.WorkRequests, error) {

	workRequests, err := r.projectRepo.GetWorkRequests(referenceId, status, page, number, document)
	if err != nil {
		return nil, err
	}
	return workRequests, nil

}
