package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Test      string   `json:"test"`
	Listening string   `json:"listening"`
	LogFile   string   `json:"log_file"`
	Database  Database `json:"database"`
}

type Database struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func Get(path string) (Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	c := Config{}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}
