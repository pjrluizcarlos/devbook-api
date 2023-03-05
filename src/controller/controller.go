package controller

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type pathVariable struct {
	r *http.Request
}

func NewPathVariable(r *http.Request) *pathVariable {
	return &pathVariable{r: r}
}

func GetAuthorizationHeader(r *http.Request) string {
	return r.Header.Get("Authorization")
}

func getRequestBody(r *http.Request) ([]byte, error) {
	requestBody, error := io.ReadAll(r.Body)
	if error != nil {
		return nil, error
	}

	return requestBody, nil
}

func (p pathVariable) uint64(key string) (uint64, error) {
	parameters := mux.Vars(p.r)
	return strconv.ParseUint(parameters["id"], 10, 64)
}
