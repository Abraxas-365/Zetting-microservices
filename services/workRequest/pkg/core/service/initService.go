package service

import (
	"errors"
	"work-request/pkg/core/models"
)

var (
	ErrUnauthorized      = errors.New("Unauthorized")
	ErrEmptyParams       = errors.New("Empty parameters")
	ErrWorkRequestExists = errors.New("WorkRequest already exists")
)

type WorkRequestService interface {
	CreateWorkRequest(workRequest models.WorkRequest) error
}

type ServiceRepository interface {
	IsUserExsist(workRequest models.WorkRequest) bool
	CreateWorkRequest(workRequest models.WorkRequest) error
}
type workRequestService struct {
	workRequestRepo ServiceRepository
}

func NewUsertService(workRequestRepo ServiceRepository) WorkRequestService {
	return &workRequestService{
		workRequestRepo,
	}
}
