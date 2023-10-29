package routes

import (
	"api/src/controllers"
	"net/http"
)

var routeLogin = Route{
	URI:          "/v1/login",
	Method:       http.MethodPost,
	Function:     controllers.Login,
	RequiredAuth: false,
}
