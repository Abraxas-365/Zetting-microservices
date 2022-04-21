package models

import "errors"

/*PROFFESSION*/
var (
	ErrDescriptionTooLong = errors.New("Description too long")
)

type Profession struct {
	Name        string `json:"name,omitempty"`
	Price       int    `json:"price,omitempty"`
	Description string `json:"description,omitempty"`
}

func (p *Profession) Validate() error {
	if len(p.Description) <= 100 {
		return ErrDescriptionTooLong
	}
	// TODO: validate mode things
	return nil
}

/*EMAIL*/
type Email string

// TODO: validate mode things
func (e Email) Validate() error {
	return nil
}

/*PASSWORD*/
type Password string

//TODO: all the function to validate
func (p Password) Validate() error {
	return nil
}
func (p Password) EqualTo(other Password) bool {
	return p == other
}
