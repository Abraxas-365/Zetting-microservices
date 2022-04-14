package service

import (
	"user/pkg/core/models"
	"user/pkg/internal/auth"
)

func (r *userService) LoginUser(email string, password string) (*models.AuthUser, error) {

	switch true {
	case email == "":
		return nil, ErrEmptyParams
	case password == "":
		return nil, ErrEmptyParams
	}

	user, _ := r.userRepo.CheckEmailExist(email)

	switch true {
	case user == nil:
		return nil, ErrUserNotFound
	case user.Password != password:
		return nil, ErrUnauthorized
	}

	token, err := auth.GereteToken(email, user.ID)
	if err != nil {
		return nil, err
	}
	return &models.AuthUser{User: *user, Token: token}, nil
}
