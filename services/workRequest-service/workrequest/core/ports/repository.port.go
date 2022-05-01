package ports

import (
	"work-request/workrequest/core/events"
	"work-request/workrequest/core/models"

	"github.com/google/uuid"
)

type WorkRequestRepository interface {
	CreateWorkRequest(workRequest models.WorkRequest) (events.Event, error)
	GetWorkRequests(referenceId uuid.UUID, status string, page int, number int, document string) (models.WorkRequests, error)
	AnswerWorkRequest(workRequest uuid.UUID, workerId uuid.UUID, status string) (events.Event, error)
	IsWorkrequestExist(newWorkRequest models.WorkRequest) bool
}
