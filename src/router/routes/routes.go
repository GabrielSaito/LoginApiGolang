package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	RequiredAuth bool
}

func Configuration(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, routeLogin)

	for _, route := range routes {
		if route.RequiredAuth {
			r.HandleFunc(route.URI, middlewares.Authentication(route.Function)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
