package controllers

import (
	"fmt"
	"log"
	"net/http"
	"teleyos/go-blog/controllers/middlewares"
)

type App struct{
	router *http.ServeMux
	server http.Server
	handler middlewares.Middleware
}

func (app *App) InitialiseRouter() {
	app.router = http.NewServeMux()
}

func (app *App) RegisterRoutes(routeList map[string]http.HandlerFunc) {
	if app.router == nil {
		log.Default().Fatalln("You should initialise the router first")
	}

	for path, handler := range routeList {
		app.router.HandleFunc(path, handler)
	}
}

func (app *App) RegisterMiddlewares(middlewareList []middlewares.Middleware) {
	app.handler = middlewares.MiddlewareStack(middlewareList)
}

// port should be preceeded by ':'
func (app *App) Serve(port string) {
	if app.router == nil {
		log.Default().Fatalln("You should initialise the app router")
	}

	var handler http.Handler
	if app.handler != nil{
		handler = app.handler(app.router)
	} else {
		handler = app.router
	}

	app.server = http.Server{
		Addr: port,
		Handler: handler,
	}

	fmt.Println("Server listening on port "+port)
	app.server.ListenAndServe()
}
