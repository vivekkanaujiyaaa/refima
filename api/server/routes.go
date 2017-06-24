package server

import (
	"github.com/PumpkinSeed/refima/api/handlers"
	"github.com/PumpkinSeed/refima/api/middleware"
	"github.com/gorilla/mux"
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

	r.HandleFunc("/", middleware.Authenticate()(s.Handlers.Root)).Methods("GET")
	r.HandleFunc("/health-check", middleware.Authenticate()(s.Handlers.HealtCheck)).Methods("GET")
	return r
}

func (s *RouteStack) getHandlers() {
	s.Handlers = handlers.NewStack()
}
