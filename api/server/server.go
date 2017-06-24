package server

import (
	"net/http"

	"github.com/PumpkinSeed/refima/config"
)

type Server struct {
	Config config.Config
}

func NewServer(c config.Config) *Server {
	s := new(Server)
	s.Config = c
	return s
}

func (s *Server) Start() {
	r := NewRouteStack()
	http.ListenAndServe(s.Config.Listening, r.getRoutes())
}
