package models

import (
	user "projects/user/core/models"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID           uuid.UUID    `bson:"_id,omitempty" json:"id"`
	Name         Name         `bson:"name" json:"name,omitempty"`
	Image        string       `bson:"image" json:"image,omitempty"`
	Description  string       `bson:"description" json:"description,omitempty"`
	Owner        uuid.UUID    `bson:"owner" json:"owner"`
	Workers      []*uuid.UUID `bson:"workers" json:"workers"`
	Color        string       `bson:"color" json:"color,omitempty"`
	DateStarted  string       `bson:"date_started" json:"date_started,omitempty"`
	DateFinished string       `bson:"date_finished" json:"date_finished,omitempty"`
	Created      time.Time    `bson:"created_at" json:"created_at,omitempty"`
	Updated      time.Time    `bson:"updated_at" json:"updated_at,omitempty"`
}

type Projects []*Project

func (p *Project) New(userId uuid.UUID) Project {
	p.ID = uuid.New()
	p.Created = time.Now()
	p.Updated = time.Now()
	p.Owner = userId
	p.Workers = []*uuid.UUID{}
	return *p
}

type LookupProject struct {
	ID           uuid.UUID  `bson:"_id,omitempty" json:"id"`
	Name         Name       `bson:"name" json:"name,omitempty"`
	Image        string     `bson:"image" json:"image,omitempty"`
	Description  string     `bson:"description" json:"description,omitempty"`
	Owner        user.User  `bson:"owner" json:"owner"`
	Workers      user.Users `bson:"workers" json:"workers"`
	Color        string     `bson:"color" json:"color,omitempty"`
	DateStarted  string     `bson:"date_started" json:"date_started,omitempty"`
	DateFinished string     `bson:"date_finished" json:"date_finished,omitempty"`
	Created      time.Time  `bson:"created_at" json:"created_at,omitempty"`
	Updated      time.Time  `bson:"updated_at" json:"updated_at,omitempty"`
}

type LookupProjects []*LookupProject
