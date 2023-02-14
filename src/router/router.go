package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func Build() *mux.Router {
	r := mux.NewRouter()
	return routes.AddRoutes(r)
}
