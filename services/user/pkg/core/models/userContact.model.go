package models

//TODO: create a validate func in servie domain
//
/*Contact*/
type Contact struct {
	Email                   Email  `json:"email"`
	Phone                   string `json:"phone,omitempty"`
	Country                 string `json:"country,omitempty"`
	IdentifierDocumentation string `json:"identifier_document,omitempty"`
}

//TODO: all the function to validate
func (c *Contact) Verified() error {
	return nil
}
