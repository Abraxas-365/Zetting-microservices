package models

import (
	"time"
)

type Project struct {
	ID           interface{} `bson:"_id,omitempty" json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Owners       interface{} `json:"owners"`
	Workers      interface{} `json:"workers"`
	Color        string      `json:"color"`
	DateStarted  string      `json:"date_started"`
	DateFinished string      `json:"date_finish"`
	Created      time.Time   `json:"created_at"`
	Updated      time.Time   `json:"updated_at,omitempty"`
}

type Projects []*Project
