package main

import (
	"teleyos/go-blog/controllers"
	"teleyos/go-blog/models/routes"
)

var App = controllers.App{}

func main() {
	App.RegisterRoutes(routes.RoutesList)

	App.Serve(":8090")
}
