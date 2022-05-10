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
	ID          uuid.UUID `bson:"_id,omitempty" json:"id"`
	Name        string    `bson:"name" json:"name,omitempty"`
	Password    Password  `bson:"password" json:"password,omitempty"`
	PerfilImage string    `bson:"perfil_image" json:"perfil_image,omitempty"`
	Gender      string    `bson:"gender" json:"gender,omitempty"`
	Age         int       `bson:"age" json:"age,omitempty"`
	Profession  string    `bson:"profession" json:"profession,omitempty"`
	Contact     Contact   `bson:"contact" json:"contact,omitempty"`
	Verified    bool      `bson:"verified" json:"verified,omitempty"`
	Created     time.Time `bson:"created_at" json:"created_at,omitempty"`
	Updated     time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}

//TODO change the user to UserPublic in create user functions
type UserPublic struct {
	ID          uuid.UUID `bson:"_id,omitempty" json:"id"`
	Name        string    `bson:"name" json:"name,omitempty"`
	PerfilImage string    `bson:"perfil_image" json:"perfil_image,omitempty"`
	Verified    bool      `bson:"verified" json:"verified,omitempty"`
	Gender      string    `bson:"gender" json:"gender,omitempty"`
	Age         int       `bson:"age" json:"age,omitempty"`
}

func (u *User) New() User {
	u.ID = uuid.New()
	u.Created = time.Now()
	u.Updated = time.Now()
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
		Gender:      u.Gender,
		Age:         u.Age,
		Verified:    u.Verified,
	}
}
