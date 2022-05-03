package application

import (
	"errors"
	project "projects/project/application"
	"projects/workrequest/core/models"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
)

type WorkRequestApplication interface {
	WorkRequestAccepted(workrequest models.WorkRequest) error
}
type workRequestApplication struct {
	projectApp project.ProjectApplication
}

func NewWorkRequestApplication(projectApp project.ProjectApplication) WorkRequestApplication {
	return &workRequestApplication{
		projectApp,
	}

}
