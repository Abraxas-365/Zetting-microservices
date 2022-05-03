package application

import (
	"work-request/workrequest/core/models"
	"work-request/workrequest/core/ports"
	"work-request/workrequest/core/service"

	"github.com/google/uuid"
)

type WorkRequestApplication interface {
	CreateWorkRequest(workRequest models.WorkRequest) error
	GetWorkRequests(referenceId uuid.UUID, status string, page int, number int, document string) (models.LookUpWorkRequests, error)
	AnswerWorkRequest(workRequest uuid.UUID, workerId uuid.UUID, status models.Status) error
}
type workRequestApplication struct {
	projectRepo    ports.WorkRequestRepository
	mqpublisher    ports.WorkRequestMQPublisher
	projectService service.WorkRequestService
}

func NewWorkRequestApplication(workRequestRepo ports.WorkRequestRepository, mqPublisher ports.WorkRequestMQPublisher, service service.WorkRequestService) WorkRequestApplication {

	return &workRequestApplication{
		workRequestRepo,
		mqPublisher,
		service,
	}

}
