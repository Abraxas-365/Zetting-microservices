package models

type AddUserToProject struct {
	OwnerId     interface{} `json:"owner_id,omitempty"`
	ProjectId   interface{} `json:"project_id,omitempty"`
	UserToAddId interface{} `json:"user_id,omitempty"`
}
