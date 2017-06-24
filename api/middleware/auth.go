package middleware

import (
	"log"
	"net/http"
)

type Adapter func(http.HandlerFunc) http.HandlerFunc

func Authenticate() Adapter {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("before")
			defer log.Println("after")
			h.ServeHTTP(w, r)
		})
	}
}
