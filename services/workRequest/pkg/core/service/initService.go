package service

import (
	"errors"
	"work-request/pkg/core/models"
	"work-request/pkg/core/ports"
)

var (
	ErrUnauthorized      = errors.New("Unauthorized")
	ErrEmptyParams       = errors.New("Empty parameters")
	ErrWorkRequestExists = errors.New("WorkRequest already exists")
)

type WorkRequestService interface {
	CreateWorkRequest(workRequest models.WorkRequest) error
}

type workRequestService struct {
	workRequestRepo ports.WorkRequestRepository
}

func NewWorRequesttService(workRequestRepo ports.WorkRequestRepository) WorkRequestService {
	return &workRequestService{
		workRequestRepo,
	}
}
