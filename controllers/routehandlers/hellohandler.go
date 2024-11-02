package routehandlers

import (
	"log"
	"net/http"
	"teleyos/go-blog/views"
)

func Hellohandler(w http.ResponseWriter, r *http.Request){
	component := views.Hello(r.PathValue("name"))
	component.Render(r.Context(),w)
	log.Default().Println(r.Header)
}
