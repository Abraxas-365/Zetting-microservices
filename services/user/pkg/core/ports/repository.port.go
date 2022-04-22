package ports

import "user/pkg/core/models"

type UserRepository interface {
	UpdateUser(query interface{}, userId interface{}) error
	CreateUser(user models.User) error
	GetUserByEmail(email models.Email) (models.User, error)
	IsUserExsist(user models.User) bool
}
