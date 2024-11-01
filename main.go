package main

import (
	"context"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	component := hello("World")
	component.Render(context.Background(), w)
}

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		os.Exit(1)
	}
}
