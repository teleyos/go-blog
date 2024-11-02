package routes

import "net/http"
import "teleyos/go-blog/controllers/routehandlers"

var RoutesList = map[string]http.HandlerFunc{
	"/": routehandlers.IndexHandler,
	"/hello/{name}": routehandlers.Hellohandler,
}
