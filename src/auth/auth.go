package auth

import (
	"devbook-api/src/config"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const userIdKey = "userId"

func GetToken(userId uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims[userIdKey] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(config.SecretKey)
}

func ValidateToken(bearerToken string) error {
	tokenString, error := extractToken(bearerToken)
	if error != nil {
		return error
	}

	token, error := jwt.Parse(tokenString, getSecretKey)
	if error != nil {
		return error
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token is invalid")
}

func GetUserId(bearerToken string) (uint64, error) {
	claims, error := getClaims(bearerToken)
	if error != nil {
		return 0, error
	}

	userId, error := strconv.ParseUint(fmt.Sprintf("%.0f", claims[userIdKey]), 10, 64)
	if error != nil {
		return 0, error
	}

	return userId, nil
}

func getClaims(bearerToken string) (jwt.MapClaims, error) {
	tokenString, error := extractToken(bearerToken)
	if error != nil {
		return nil, error
	}

	token, error := jwt.Parse(tokenString, getSecretKey)
	if error != nil {
		return nil, error
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token is invalid")
}

func extractToken(bearerToken string) (string, error) {
	splitedBearerToken := strings.Split(bearerToken, " ")

	if len(splitedBearerToken) == 2 {
		return splitedBearerToken[1], nil
	}

	return "", errors.New("bearer token is invalid")
}

func getSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("wrong assignature method: [%v]", token.Header["alg"])
	}

	return config.SecretKey, nil
}
