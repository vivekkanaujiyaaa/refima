package api

import (
	"github.com/PumpkinSeed/refima/api/server"
	"github.com/PumpkinSeed/refima/config"
)

func Start(c config.Config) {
	s := server.New(c)
	s.Start()
}
