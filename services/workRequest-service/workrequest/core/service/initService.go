package service

import (
	"errors"
	"work-request/workrequest/core/models"
	"work-request/workrequest/core/ports"
)

var (
	ErrUnauthorized      = errors.New("Unauthorized")
	ErrEmptyParams       = errors.New("Empty parameters")
	ErrWorkRequestExists = errors.New("WorkRequest already exists")
)

type WorkRequestService interface {
	CanCreateWorkRequest(workRequest models.WorkRequest) error
}

type workRequestService struct {
	workRequestRepo ports.WorkRequestRepository
}

func NewWorRequesttService(workRequestRepo ports.WorkRequestRepository) WorkRequestService {
	return &workRequestService{
		workRequestRepo,
	}
}
