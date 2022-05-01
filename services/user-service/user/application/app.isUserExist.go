package application

import (
	"user/user/core/models"
)

func (a userApplication) IsUserExsist(email models.Email) error {
	if a.userRepo.IsUserExsist(email) {
		return ErrUserExists
	}
	return nil
}
