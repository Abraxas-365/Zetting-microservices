package ports

import "user/pkg/core/models"

type UserService interface {
	CreateUser(user models.User) (*models.User, error)
	LoginUser(email string, password string) (*models.AuthUser, error)
	CheckEmailExist(email string) bool
}
