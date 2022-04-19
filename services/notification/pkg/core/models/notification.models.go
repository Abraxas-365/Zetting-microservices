package models

import "time"

type Notification struct {
	ID           interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	NotifiedUser User        `bson:"notified" json:"notified,omitempty"`
	NotifierUser User        `bson:"notifier" json:"notifier,omitempty"`
	WorkRequest  WorkRequest `bson:"workrequest,omitempty" json:"workRequest,omitempty"`
	Type         string      `bson:"type" json:"type,omitempty"`
	Read         bool        `bson:"is_read"  json:"is_read,omitempty"`
	Message      string      `bson:"message" json:"message,omitempty"`
	Created      time.Time   `bson:"created_at" json:"created_at,omitempty"`
	Updated      time.Time   `bson:"updated_at" json:"updated_at,omitempty"`
}

type Notifications []*Notification

type User struct {
	ID    interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string      `bson:"name" json:"name,omitempty"`
	Image string      `bson:"image" json:"image,omitempty"`
}

type Project struct {
	ID    interface{} `json:"id,omitempty"`
	Name  string      `json:"name,omitempty"`
	Image string      `json:"image,omitempty"`
}

type WorkRequest struct {
	ID      interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Project Project     `bson:"project" json:"project,omitempty"`
	Created time.Time   `bson:"created_at" json:"created_at,omitempty"`
}
