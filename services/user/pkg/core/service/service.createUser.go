package service

import "user/pkg/core/models"

func (s *userService) CreateUser(user models.User) error {
	//TODO: applay the logic of chaeck email , check pass , etc
	if err := s.userRepo.CreateUser(user); err != nil {
		return ErrUserExists
	}
	return nil
}
