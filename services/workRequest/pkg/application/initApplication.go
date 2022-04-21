package application

import (
	"work-request/pkg/core/models"
	"work-request/pkg/core/ports"
)

type WorkRequestApplication interface {
	CreateWorkRequest(workRequest models.WorkRequest) (models.WorkRequest, error)
	GetWorkRequests(referenceId interface{}, status string, page int, number int, document string) (models.WorkRequests, error)
	AnswerWorkRequest(workRequest models.WorkRequest) error
}
type workRequestApplication struct {
	projectRepo ports.WorkRequestRepository
	mqpublisher ports.WorkRequestMQPublisher
}

func NewWorkRequestApplication(workRequestRepo ports.WorkRequestRepository, mqPublisher ports.WorkRequestMQPublisher) WorkRequestApplication {

	return &workRequestApplication{
		workRequestRepo,
		mqPublisher,
	}

}
