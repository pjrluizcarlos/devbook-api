package middleware

import (
	"devbook-api/src/auth"
	"devbook-api/src/response"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")

		if error := auth.ValidateToken(bearerToken); error != nil {
			response.Error(w, http.StatusUnauthorized, error)
		} else {
			next(w, r)
		}
	}
}
