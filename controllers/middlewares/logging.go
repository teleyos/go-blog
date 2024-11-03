package middlewares

import (
	"log"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Default().Println(r.URL)
	})
}
