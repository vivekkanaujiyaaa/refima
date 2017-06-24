package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/PumpkinSeed/refima/api/auth"
)

type Adapter func(http.HandlerFunc) http.HandlerFunc

func Authenticate() Adapter {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if !auth.Check(authHeader) {
				Unauthorized(w, r)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	s := Error{
		Message: "unauthorized",
		Status:  http.StatusUnauthorized,
	}
	resp, _ := json.Marshal(s)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(resp)
}
