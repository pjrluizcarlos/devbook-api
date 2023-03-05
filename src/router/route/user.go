package route

import (
	"devbook-api/src/controller"
	"net/http"
)

const userResource = "/users"
const userResourceByID = userResource + "/{id}"

var userRoutes = []Route{
	{
		URI:             userResource,
		Method:          http.MethodGet,
		Handler:         controller.FindAllUsers,
		IsAuthenticated: true,
	},
	{
		URI:             userResourceByID,
		Method:          http.MethodGet,
		Handler:         controller.FindUserByID,
		IsAuthenticated: true,
	},
	{
		URI:             userResource,
		Method:          http.MethodPost,
		Handler:         controller.CreateUser,
		IsAuthenticated: false,
	},
	{
		URI:             userResourceByID,
		Method:          http.MethodDelete,
		Handler:         controller.DeleteUserByID,
		IsAuthenticated: true,
	},
	{
		URI:             userResourceByID,
		Method:          http.MethodPut,
		Handler:         controller.UpdateUserByID,
		IsAuthenticated: true,
	},
	{
		URI:             userResourceByID + "/follow",
		Method:          http.MethodPost,
		Handler:         controller.FollowUser,
		IsAuthenticated: true,
	},
	{
		URI:             userResourceByID + "/unfollow",
		Method:          http.MethodPost,
		Handler:         controller.UnfollowUser,
		IsAuthenticated: true,
	},
	{
		URI:             userResourceByID + "/followers",
		Method:          http.MethodGet,
		Handler:         controller.FindAllFollowersById,
		IsAuthenticated: true,
	},
}
