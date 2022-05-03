package models

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProjectNameTooLong  = errors.New("Project name too long")
	ErrProjectNameTooShort = errors.New("Project name too short")
	ErrNoProjectID         = errors.New("No Project ID")
)

type Project struct {
	ID    uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string    `bson:"name" json:"name,omitempty"`
	Image string    `bson:"image" json:"image,omitempty"`
}

type Projects []*Project
