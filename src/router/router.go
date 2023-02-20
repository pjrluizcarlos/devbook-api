package router

import (
	"devbook-api/src/router/route"

	"github.com/gorilla/mux"
)

func Build() *mux.Router {
	r := mux.NewRouter()
	return route.AddRoutes(r)
}
