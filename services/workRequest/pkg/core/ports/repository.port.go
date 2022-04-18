package ports

import "work-request/pkg/core/models"

type WorkRequestRepository interface {
	CreateWorkRequest(workRequest models.WorkRequest) (*models.WorkRequest, error)
	GetWorkRequests(referenceId interface{}, status string, page int, number int, document string) (models.WorkRequests, error)
	AnswerWorkRequest(workRequest models.WorkRequest) error
	UpdateWorkRequest(query interface{}, workRequestId interface{}) error
}
