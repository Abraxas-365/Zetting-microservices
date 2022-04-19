package models

type Contact struct {
	Email                   string `json:"email"`
	Phone                   string `json:"phone,omitempty"`
	Country                 string `json:"country,omitempty"`
	IdentifierDocumentation string `json:"identifier_document,omitempty"`
}
