package config

import "testing"

func TestGet(t *testing.T) {
	conf, _ := Get("../config.json")
	if conf.Test != "test" {
		t.Error(`Test config should be "test"`)
	}
}
