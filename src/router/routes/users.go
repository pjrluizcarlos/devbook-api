package routes

import (
	"api/src/controllers"
	"net/http"
)

const userResource = "/users"
const userResourceByID = userResource + "/{id}"

var userRoutes = []Route{
	{
		URI:             userResource,
		Method:          http.MethodGet,
		Handler:         controllers.FindAllUsers,
		IsAuthenticated: false,
	},
	{
		URI:             userResourceByID,
		Method:          http.MethodGet,
		Handler:         controllers.FindUserByID,
		IsAuthenticated: false,
	},
	{
		URI:             userResource,
		Method:          http.MethodPost,
		Handler:         controllers.CreateUser,
		IsAuthenticated: false,
	},
	{
		URI:             userResourceByID,
		Method:          http.MethodDelete,
		Handler:         controllers.DeleteUserByID,
		IsAuthenticated: false,
	},
	{
		URI:             userResourceByID,
		Method:          http.MethodPut,
		Handler:         controllers.UpdateUserByID,
		IsAuthenticated: false,
	},
}
