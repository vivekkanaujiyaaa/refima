package server

import "github.com/gorilla/mux"
import "github.com/PumpkinSeed/refima/api/handlers"

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
	r.HandleFunc("/", s.Handlers.Root).Methods("GET")
	r.HandleFunc("/health-check", s.Handlers.HealtCheck).Methods("GET")
	return r
}

func (s *RouteStack) getHandlers() {
	s.Handlers = handlers.NewStack()
}
