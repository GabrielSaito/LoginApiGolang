package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/v1/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		RequiredAuth: false,
	},
	{
		URI:          "/v1/users",
		Method:       http.MethodGet,
		Function:     controllers.SearchUsers,
		RequiredAuth: true,
	},
	{
		URI:          "/v1/users/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.SearchUserId,
		RequiredAuth: true,
	},
	{
		URI:          "/v1/user/{userId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		RequiredAuth: true,
	},
	{
		URI:          "/v1/users/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		RequiredAuth: true,
	},
	{
		URI:          "/v1/users/{userId}/resetPassword",
		Method:       http.MethodPost,
		Function:     controllers.ResetPassword,
		RequiredAuth: true,
	},
}
