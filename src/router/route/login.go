package route

import (
	"devbook-api/src/controller"
	"net/http"
)

var loginRoute = Route{
	URI:             "/login",
	Method:          http.MethodPost,
	Handler:         controller.Login,
	IsAuthenticated: false,
}
