package config

import "testing"

func TestGet(t *testing.T) {
	conf, err := Get("../config.json")
	if err != nil {
		t.Errorf(`Config error should be nil, instead of %s`, err.Error())
	}
	if conf.Test != "test" {
		t.Error(`Test config should be "test"`)
	}
}
