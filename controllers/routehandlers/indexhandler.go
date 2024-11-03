package routehandlers

import (
	"net/http"
	"teleyos/go-blog/views"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	views.Index().Render(r.Context(),w)
}
