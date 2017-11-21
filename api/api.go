package api

import (
	logging "github.com/sirupsen/logrus"
	"github.com/xalabs/refima/api/server"
	"github.com/xalabs/refima/config"
)

func Start(c config.Config, l *logging.Entry) {
	s := server.New(c, l)
	s.Start()
}
