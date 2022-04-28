package models

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUserNameTooLong  = errors.New("User name too long")
	ErrUserNameTooShort = errors.New("User name too short")
	ErrNoUserID         = errors.New("No user ID")
)

type User struct {
	ID    uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string    `bson:"name" json:"name,omitempty"`
	Image string    `bson:"perfil_image" json:"perfil_image,omitempty"`
}
