package routehandlers

import (
	"net/http"
	"teleyos/go-blog/views"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	component := views.Hello("index")
	component.Render(r.Context(), w)
}
