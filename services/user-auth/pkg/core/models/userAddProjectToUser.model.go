package models

type AddProjectToUser struct {
	Owner   interface{} `json:"owner"` //esto va a venir de el work notification
	User    interface{} `json:"user"`
	Project interface{} `json:"project"`
}
