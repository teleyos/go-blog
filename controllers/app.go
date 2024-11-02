package controllers

import (
	"os"
	"net/http"
)

type App struct{
}

func (app *App) RegisterRoutes(routeList map[string]http.HandlerFunc) {
	for path, handler := range routeList {
		http.HandleFunc(path, handler)
	}
}

func (app *App) Serve(port string) {
	err := http.ListenAndServe(port, nil)
	if err != nil {
		os.Exit(1)
	}
}
