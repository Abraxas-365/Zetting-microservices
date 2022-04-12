package models

import "time"

type Notification struct {
	ID                interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	NotifiedUserId    interface{} `bson:"notified_id" json:"notified_id,omitempty"`
	NotifierUserImage string      `bson:"notifier_image" json:"notifier_image,omitempty"`
	NotifierUserId    interface{} `bson:"notifier_id" json:"notifier_id,omitempty"`
	ProjectName       string      `bson:"project_id,omitempty" json:"project_id,omitempty"`
	Type              string      `bson:"type" json:"type,omitempty"`
	ReferenceId       interface{} `bson:"reference_id" json:"reference_id,omitempty"`
	Read              bool        `bson:"is_read"  json:"is_read,omitempty"`
	Message           string      `bson:"message" json:"message,omitempty"`
	Created           time.Time   `bson:"created_at" json:"created_at,omitempty"`
	Updated           time.Time   `bson:"updated_at" json:"updated_at,omitempty"`
}

type Notifications []*Notification
