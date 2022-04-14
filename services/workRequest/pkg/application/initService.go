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
	mqpublisher ports.WorkRequestMQueue
}

func NewWorkRequestService(workRequestRepo ports.WorkRequestRepository, mqPublisher ports.WorkRequestMQueue) ports.WorkRequestService {
	return &workRequestService{
		workRequestRepo,
		mqPublisher,
	}

}
