package route

import (
	"devbook-api/src/controller"
	"net/http"
)

const postResource = "/posts"
const postResourceByID = postResource + "/{id}"

var postRoutes = []Route{
	{
		URI:             postResource,
		Method:          http.MethodGet,
		Handler:         controller.FindAllPosts,
		IsAuthenticated: true,
	},
	{
		URI:             postResourceByID,
		Method:          http.MethodGet,
		Handler:         controller.FindPostByID,
		IsAuthenticated: true,
	},
	{
		URI:             postResource,
		Method:          http.MethodPost,
		Handler:         controller.CreatePost,
		IsAuthenticated: true,
	},
	{
		URI:             postResourceByID,
		Method:          http.MethodDelete,
		Handler:         controller.DeletePostByID,
		IsAuthenticated: true,
	},
	{
		URI:             postResourceByID,
		Method:          http.MethodPut,
		Handler:         controller.UpdatePostByID,
		IsAuthenticated: true,
	},
}
