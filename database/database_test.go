package database

import (
	"fmt"
	"testing"

	"github.com/PumpkinSeed/refima/config"
)

var operation Operation

func TestOpenDB(t *testing.T) {
	conf := config.Config{
		Database: config.Database{
			Connection: "./test.db",
			Dialect:    "sqlite3",
		},
	}
	var err error
	operation, err = NewOperationHandler(conf)
	if err != nil {
		t.Errorf(`Error should be nil, instead of %s`, err.Error())
	}
}

func TestNewUser(t *testing.T) {
	err := operation.NewUser("test", "test")
	if err != nil {
		t.Errorf(`Error should be nil, instead of %s`, err.Error())
	}
}

func TestGetUser(t *testing.T) {
	user := User{
		Name: "test",
	}
	result := operation.GetUser(user)
	fmt.Println(result)
}
