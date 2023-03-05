package controller

import (
	"io"
	"net/http"
)

func getRequestBody(r *http.Request) ([]byte, error) {
	requestBody, error := io.ReadAll(r.Body)
	if error != nil {
		return nil, error
	}

	return requestBody, nil
}

func GetAuthorizationHeader(r *http.Request) string {
	return r.Header.Get("Authorization")
}
