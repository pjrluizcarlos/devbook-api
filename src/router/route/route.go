package route

import (
	"devbook-api/src/middleware"
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
	var routes []Route

	routes = append(routes, userRoutes...)
	routes = append(routes, postRoutes...)
	routes = append(routes, loginRoute)

	fmt.Println("Routes creation started.")

	for _, route := range routes {
		addRoute(r, route)
	}

	fmt.Println("Routes creation completed.")

	return r
}

func addRoute(r *mux.Router, route Route) {
	r.HandleFunc(route.URI, getHandleFunc(route)).Methods(route.Method)
	fmt.Printf("Route [%s]	%s added.\n", route.Method, route.URI)
}

func getHandleFunc(route Route) http.HandlerFunc {
	if route.IsAuthenticated {
		return middleware.Authenticate(route.Handler)
	}

	return route.Handler
}
