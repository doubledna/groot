package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user_id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("groottoken"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
