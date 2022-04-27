package auth

import "github.com/golang-jwt/jwt/v4"

func GereteToken(email string, id interface{}) (string, error) {

	secret := "JWT_SECRET_KEY"
	claims := jwt.MapClaims{
		"email": email,
		"id":    id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
