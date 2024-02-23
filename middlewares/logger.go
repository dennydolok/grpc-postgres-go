package middlewares

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request URI", r.RequestURI)
		log.Println("Request METHOD :", r.Method)
		log.Println("Request HOST : ", r.Host)

		next.ServeHTTP(w, r)
	})
}
