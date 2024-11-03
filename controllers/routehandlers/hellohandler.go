package routehandlers

import (
	"net/http"
	"teleyos/go-blog/views"
)

func Hellohandler(w http.ResponseWriter, r *http.Request){
	views.Hello(r.URL.Query().Get("name")).Render(r.Context(),w)
}
