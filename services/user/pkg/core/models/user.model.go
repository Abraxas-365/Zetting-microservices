package models

import (
	"errors"
	"time"
)

var (
	ErrUserNotFound    = errors.New("User not found")
	ErrEmptyParams     = errors.New("Empty parameters")
	ErrInvalidPassword = errors.New("Invalid password")
)

type User struct {
	ID          interface{} `bson:"_id,omitempty" json:"id"`
	Name        string      `bson:"name" json:"name,omitempty"`
	Password    Password    `bson:"password" json:"password,omitempty"`
	PerfilImage string      `bson:"perfil_image" json:"perfil_image,omitempty"`
	Contact     Contact     `bson:"contact" json:"contact,omitempty"`
	Profession  Profession  `bson:"profession" json:"profession,omitempty"`
	Features    Features    `bson:"features" json:"features,omitempty"`
	Verified    bool        `bson:"verified" json:"verified,omitempty"`
	Created     time.Time   `bson:"created_at" json:"created_at,omitempty"`
	Updated     time.Time   `bson:"updated_at" json:"updated_at,omitempty"`
}

func (u *User) Validate() error {
	switch true {
	case u.Contact.Email == "":
		return ErrEmptyParams
	case u.Password == "":
		return ErrEmptyParams
	}
	// TODO: validate
	return nil
}

func (u *User) ExposeToPublic() {
	u.Password = ""
	u.Contact = Contact{}
}
