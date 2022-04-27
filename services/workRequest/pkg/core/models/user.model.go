package models

import "errors"

var (
	ErrUserNameTooLong  = errors.New("User name too long")
	ErrUserNameTooShort = errors.New("User name too short")
	ErrNoUserID         = errors.New("No user ID")
)

type User struct {
	ID    interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string      `bson:"name" json:"name,omitempty"`
	Image string      `bson:"perfil_image" json:"perfil_image,omitempty"`
}

func (u *User) Validate() error {
	switch len := len(u.Name); {
	case len <= 0:
		return ErrUserNameTooShort
	case len > 20:
		return ErrUserNameTooLong
	}
	switch u.ID {
	case "":
		return ErrNoUserID
	case 0:
		return ErrNoUserID
	case nil:
		return ErrNoUserID
	}
	return nil
}
