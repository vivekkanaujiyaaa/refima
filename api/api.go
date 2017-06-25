package api

import (
	"github.com/PumpkinSeed/refima/api/server"
	"github.com/PumpkinSeed/refima/config"
	logging "github.com/sirupsen/logrus"
)

func Start(c config.Config, l *logging.Entry) {
	s := server.New(c, l)
	s.Start()
}
