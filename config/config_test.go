package config

import "testing"

func TestGet(t *testing.T) {
	conf, err := Get("../config.json")
	if err != nil {
		t.Error(`Config error should be nil`)
	}
	if conf.Test != "test" {
		t.Error(`Test config should be "test"`)
	}
}
