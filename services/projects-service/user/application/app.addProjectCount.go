package application

import "github.com/google/uuid"

func (a *userApplication) AddProjectCount(userId uuid.UUID) error {
	if err := a.repo.AddProjectCount(userId); err != nil {
		//TODO do something with the error
		return nil
	}
	return nil
}
