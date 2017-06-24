package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Test      string `json:"test"`
	Listening string `json:"listening"`
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
