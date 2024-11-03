package main

import (
	"teleyos/go-blog/controllers"
	"teleyos/go-blog/controllers/middlewares"
	"teleyos/go-blog/models/routes"
)

var App = controllers.App{}

func main() {

	App.InitialiseRouter()

	App.RegisterRoutes(routes.RoutesList)

	App.RegisterMiddlewares(middlewares.MiddlewareList)

	App.Serve(":8090")
}
