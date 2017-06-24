package server

import (
	"net/http"
	"os"

	"github.com/PumpkinSeed/refima/config"
	"github.com/labstack/gommon/log"
	logging "github.com/op/go-logging"
)

type Server struct {
	Config config.Config
	Log    *logging.Logger
}

func New(c config.Config, l *logging.Logger) *Server {
	s := new(Server)
	s.Config = c
	s.Log = l
	return s
}

func (s *Server) Start() {
	r := NewRouteStack()
	s.Log.Infof("Server listening on: %s", s.Config.Listening)
	err := http.ListenAndServe(s.Config.Listening, r.getRoutes())
	if err != nil {
		log.Errorf("Server start failed -> %s", err.Error())
		os.Exit(0)
	}
}
