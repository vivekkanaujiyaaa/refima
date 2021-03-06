package server

import (
	"github.com/gorilla/mux"
	"github.com/xalabs/refima/api/handlers"
	"github.com/xalabs/refima/api/middleware"
)

type RouteStack struct {
	Handlers *handlers.Stack
}

func NewRouteStack() *RouteStack {
	s := new(RouteStack)
	s.getHandlers()
	return s
}

func (s *RouteStack) getRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", s.Handlers.Login).Methods("POST")

	r.HandleFunc("/", middleware.Authenticate()(s.Handlers.Root)).Methods("GET")
	r.HandleFunc("/health-check", middleware.Authenticate()(s.Handlers.HealtCheck)).Methods("GET")
	r.HandleFunc("/logout", middleware.Authenticate()(s.Handlers.Logout)).Methods("GET")
	return r
}

func (s *RouteStack) getHandlers() {
	s.Handlers = handlers.NewStack()
}
