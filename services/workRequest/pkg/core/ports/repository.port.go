package ports

import (
	"work-request/pkg/core/events"
	"work-request/pkg/core/models"
)

type WorkRequestRepository interface {
	CreateWorkRequest(workRequest models.WorkRequest) (events.Event, error)
	GetWorkRequests(referenceId interface{}, status string, page int, number int, document string) (models.WorkRequests, error)
	AnswerWorkRequest(workRequest models.WorkRequest) (events.Event, error)
	IsWorkrequestExist(newWorkRequest models.WorkRequest) bool
}
