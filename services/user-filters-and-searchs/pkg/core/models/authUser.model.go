package models

type AuthUser struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
