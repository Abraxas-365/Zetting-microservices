package ports

import "user/pkg/core/models"

type UserRepository interface {
	CreateUser(user models.User) (*models.User, error)
	CheckEmailExist(email string) (*models.User, error)
}
