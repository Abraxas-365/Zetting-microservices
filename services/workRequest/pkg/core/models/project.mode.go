package models

import "errors"

var (
	ErrProjectNameTooLong  = errors.New("Project name too long")
	ErrProjectNameTooShort = errors.New("Project name too short")
	ErrNoProjectID         = errors.New("No Project ID")
)

type Project struct {
	ID    interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string      `bson:"name" json:"name,omitempty"`
	Image string      `bson:"image" json:"image,omitempty"`
}

func (p *Project) Validate() error {
	switch len := len(p.Name); {
	case len <= 0:
		return ErrProjectNameTooShort
	case len > 20:
		return ErrProjectNameTooLong
	}

	switch p.ID {
	case "":
		return ErrNoProjectID
	case 0:
		return ErrNoProjectID
	case nil:
		return ErrNoProjectID
	}
	return nil
}
