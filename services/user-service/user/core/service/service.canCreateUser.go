package service

import (
	"user/user/core/models"
)

func (s *userService) CanCreateUser(user models.User) error {
	//TODO: applay the logic of chaeck email , check pass , etc
	if s.userRepo.IsUserExsist(user.Contact.Email) {
		return ErrUserExists
	}
	return nil
}
