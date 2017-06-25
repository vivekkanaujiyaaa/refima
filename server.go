package main

import (
	"fmt"
	"os"

	"github.com/PumpkinSeed/refima/api"
	"github.com/PumpkinSeed/refima/config"
	"github.com/PumpkinSeed/refima/database"
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
		{
			Name:  "migrate",
			Usage: "Migrating the database",
			Action: func(c *cli.Context) error {
				return Migrate(path)
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
	log := GetLogger()
	log.Info("Refima - Remote file manager")
	conf, err := config.Get(configPath)
	if err != nil {
		log.Errorf("Load config file failed -> %s", err.Error())
		os.Exit(0)
		return err
	}
	api.Start(conf, log)
	return nil
}

func Migrate(configPath string) error {
	log := GetLogger()
	log.Info("Refima - Remote file manager")
	conf, err := config.Get(configPath)
	if err != nil {
		log.Errorf("Load config file failed -> %s", err.Error())
		os.Exit(0)
		return err
	}
	err = database.Migrate(conf)
	if err != nil {
		log.Errorf("Load config file failed -> %s", err.Error())
		os.Exit(0)
		return err
	}
	return nil
}

func GetLogger() *logging.Entry {
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
