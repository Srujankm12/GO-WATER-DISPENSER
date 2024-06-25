package middlewares

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter ,r *http.Request){
		log.Printf("The Request Type Is %v on Route %v" , r.Method , r.URL)
		next.ServeHTTP(w,r)
	})
}

