package models

type AddUserToProject struct {
	Owner     interface{} `json:"owner,omitempty"`
	Project   interface{} `json:"project,omitempty"`
	UserToAdd interface{} `json:"user,omitempty"`
}
