package main

import (
	"fmt"
	"os"

	"github.com/PumpkinSeed/refima/api/server"
	"github.com/PumpkinSeed/refima/config"
	logging "github.com/sirupsen/logrus"
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

func getLogger() *logging.Entry {
	f, err := os.OpenFile("refima.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	logging.SetFormatter(&logging.JSONFormatter{})
	logging.SetOutput(f)
	logging.SetLevel(logging.InfoLevel)

	return logging.WithFields(logging.Fields{
		"service": "refima",
	})
}
