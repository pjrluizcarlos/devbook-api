package auth

import (
	"devbook-api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GetToken(userId uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(config.SecretKey)
}
