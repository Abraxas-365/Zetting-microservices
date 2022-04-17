package service

import (
	"errors"
	"work-request/pkg/core/ports"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

type workRequestService struct {
	projectRepo ports.WorkRequestRepository
	mqpublisher ports.WorkRequestMQPublisher
}

func NewWorkRequestService(workRequestRepo ports.WorkRequestRepository, mqPublisher ports.WorkRequestMQPublisher) ports.WorkRequestService {
	return &workRequestService{
		workRequestRepo,
		mqPublisher,
	}

}
