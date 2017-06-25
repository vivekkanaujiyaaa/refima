package database

import (
	"testing"

	"github.com/PumpkinSeed/refima/config"
	"github.com/PumpkinSeed/tuid"
	mocket "github.com/Selvatico/go-mocket"
)

var operation Operation

func init() {
	mocket.Catcher.Register()
	conf := config.Config{
		Database: config.Database{
			Connection: "RANDOM_STRING",
			Dialect:    mocket.DRIVER_NAME,
		},
	}
	operation, _ = NewOperationHandler(conf)
}

func TestNewUser(t *testing.T) {
	g := tuid.NewGenerator(5, true, false)
	id, _ := g.New()
	user := User{
		ID:       id,
		Name:     "test",
		Password: "test",
	}
	operation.NewUser(user)
}
