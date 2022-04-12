package models

type AddProjectToUser struct {
	OwnerId   interface{} `json:"owner"` //esto va a venir de el work notification
	UserId    interface{} `json:"user"`
	ProjectId interface{} `json:"project"`
}
