package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI             string
	Method          string
	Handler         func(http.ResponseWriter, *http.Request)
	IsAuthenticated bool
}

func AddRoutes(r *mux.Router) *mux.Router {
	routes := userRoutes

	fmt.Println("Routes creation started.")

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
		fmt.Printf("Route [%s]	%s created.\n", route.Method, route.URI)
	}

	fmt.Println("Routes creation completed.")

	return r
}
