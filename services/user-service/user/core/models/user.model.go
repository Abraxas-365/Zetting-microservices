package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound    = errors.New("User not found")
	ErrEmptyParams     = errors.New("Empty parameters")
	ErrInvalidPassword = errors.New("Invalid password")
)

type User struct {
	ID          uuid.UUID  `bson:"_id,omitempty" json:"id"`
	Name        string     `bson:"name" json:"name,omitempty"`
	Password    Password   `bson:"password" json:"password,omitempty"`
	PerfilImage string     `bson:"perfil_image" json:"perfil_image,omitempty"`
	Contact     Contact    `bson:"contact" json:"contact,omitempty"`
	Profession  Profession `bson:"profession" json:"profession,omitempty"`
	Features    Features   `bson:"features" json:"features,omitempty"`
	Verified    bool       `bson:"verified" json:"verified,omitempty"`
	Created     time.Time  `bson:"created_at" json:"created_at,omitempty"`
	Updated     time.Time  `bson:"updated_at" json:"updated_at,omitempty"`
}

//TODO change the user to UserPublic in create user functions
type UserPublic struct {
	ID          uuid.UUID  `bson:"_id,omitempty" json:"id"`
	Name        string     `bson:"name" json:"name,omitempty"`
	PerfilImage string     `bson:"perfil_image" json:"perfil_image,omitempty"`
	Profession  Profession `bson:"profession" json:"profession,omitempty"`
	Features    Features   `bson:"features" json:"features,omitempty"`
	Verified    bool       `bson:"verified" json:"verified,omitempty"`
}

func (u *User) New() User {
	u.ID = uuid.New()
	u.Created = time.Now()
	u.Updated = time.Now()
	u.Features.Skills = []string{}
	u.Verified = false
	u.PerfilImage = "/noPerfil.png"
	return *u
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

func (u *User) ExposeToPublic() UserPublic {
	return UserPublic{
		ID:          u.ID,
		Name:        u.Name,
		PerfilImage: u.PerfilImage,
		Profession:  u.Profession,
		Features:    u.Features,
		Verified:    u.Verified,
	}
}
