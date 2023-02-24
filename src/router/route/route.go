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
	routes := append(userRoutes, loginRoute)

	fmt.Println("Routes creation started.")

	for _, route := range routes {
		var handlerFunc http.HandlerFunc

		if route.IsAuthenticated {
			handlerFunc = middleware.Authenticate(route.Handler)
		} else {
			handlerFunc = route.Handler
		}

		r.HandleFunc(route.URI, handlerFunc).Methods(route.Method)

		fmt.Printf("Route [%s]	%s created.\n", route.Method, route.URI)
	}

	fmt.Println("Routes creation completed.")

	return r
}
