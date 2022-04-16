package models

import (
	"time"
)

type Project struct {
	ID           interface{} `bson:"_id,omitempty" json:"id"`
	Name         string      `bson:"name" json:"name,omitempty"`
	Image        string      `bson:"image" json:"image,omitempty"`
	Description  string      `bson:"description" json:"description,omitempty"`
	Owners       interface{} `bson:"owners" json:"owners"`
	Workers      interface{} `bson:"workers" json:"workers"`
	Color        string      `bson:"color" json:"color,omitempty"`
	DateStarted  string      `bson:"date_started" json:"date_started,omitempty"`
	DateFinished string      `bson:"date_finished" json:"date_finished,omitempty"`
	Created      time.Time   `bson:"created_at" json:"created_at,omitempty"`
	Updated      time.Time   `bson:"updated_at" json:"updated_at,omitempty"`
}

type Projects []*Project

type User struct {
	ID    interface{} `bson:"id" json:"id,omitempty"`
	Name  string      `bson:"name" json:"name,omitempty"`
	Image string      `bson:"image" json:"image,omitempty"`
}
