package models

import "errors"

/*PROFFESSION*/
var (
	ErrDescriptionTooLong = errors.New("Description too long")
)

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
