package models

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `bson:"_id" json:"id,omitempty"`
	Name         string    `bson:"name" json:"name,omitempty"`
	Image        string    `bson:"perfil_image" json:"perfil_image,omitempty"`
	ProjectCount int       `bson:"projects_count" json:"projects_count"`
}

type Users []*User
