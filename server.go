package main

import (
	"os"

	"github.com/PumpkinSeed/refima/api/server"
	"github.com/PumpkinSeed/refima/config"
	logging "github.com/op/go-logging"
	"github.com/urfave/cli"
)

var path string

func main() {
	app := cli.NewApp()
	app.Name = "refima"
	app.Usage = "Remote file manager"
	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Start server",
			Action: func(c *cli.Context) error {
				return RunServer(path)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config, c",
					Value:       "./config.json",
					Usage:       "Load configuration from `FILE`",
					Destination: &path,
				},
			},
		},
	}

	app.Run(os.Args)
}

func RunServer(configPath string) error {
	log := getLogger()
	log.Info("Refima - Remote file manager")
	conf, err := config.Get(configPath)
	if err != nil {
		log.Errorf("Load config file failed -> %s", err.Error())
		os.Exit(0)
		return err
	}
	s := server.New(conf, log)
	s.Start()
	return nil
}

func getLogger() *logging.Logger {
	var log = logging.MustGetLogger("refima")
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)

	return log
}
