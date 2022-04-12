package models

type Profession struct {
	Name        string `json:"name,omitempty"`
	Price       int    `json:"price,omitempty"`
	Description string `json:"description,omitempty"`
}
