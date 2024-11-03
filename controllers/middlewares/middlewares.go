package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

var MiddlewareList = []Middleware{
	Logging,
}

func MiddlewareStack(middlewareList []Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewareList)-1; i >= 0; i-- {
			currentMiddleware := middlewareList[i]
			next = currentMiddleware(next)
		}

		return next
	}
}
