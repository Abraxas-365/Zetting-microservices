package application

import (
	"work-request/pkg/core/models"
	"work-request/pkg/core/ports"
	"work-request/pkg/core/service"

	"github.com/google/uuid"
)

type WorkRequestApplication interface {
	CreateWorkRequest(workRequest models.WorkRequest) error
	GetWorkRequests(referenceId uuid.UUID, status string, page int, number int, document string) (models.WorkRequests, error)
	AnswerWorkRequest(workRequest models.WorkRequest) error
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
