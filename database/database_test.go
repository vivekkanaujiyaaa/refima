package database

import (
	"fmt"
	"os"
	"testing"

	"github.com/PumpkinSeed/refima/config"
)

var operation Operation
var conf config.Config

func init() {
	conf = config.Config{
		Database: config.Database{
			Connection: "./test.db",
			Dialect:    "sqlite3",
		},
	}
	_, err := os.Stat(conf.Database.Connection)
	if !os.IsNotExist(err) {
		os.Remove(conf.Database.Connection)
	}
}

func TestOpenDB(t *testing.T) {
	var err error
	operation, err = NewOperationHandler(conf)
	if err != nil {
		t.Errorf(`Error should be nil, instead of %s`, err.Error())
	}
}

func TestMigration(t *testing.T) {
	err := Migrate(conf)
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
	result, err := operation.GetUser(user)
	if err != nil {
		t.Errorf(`Error should be nil, instead of %s`, err.Error())
	}
	if len(result.ID) < 20 {
		t.Errorf(`The userID should be bigger than 20, instead of %d`, len(result.ID))
	}
	fmt.Println(result)
}

func TestUpdateUser(t *testing.T) {
	user := User{
		Name: "test",
	}
	result, err := operation.GetUser(user)
	if err != nil {
		t.Errorf(`Error should be nil, instead of %s`, err.Error())
	}
	result.Name = "after_test"
	result.Password = "after_test"
	operation.UpdateUser(result)
	userAfterUpdate := User{
		ID: result.ID,
	}
	resultAfterUpdate, err := operation.GetUser(userAfterUpdate)
	if err != nil {
		t.Errorf(`Error should be nil, instead of %s`, err.Error())
	}
	if resultAfterUpdate.Name != "after_test" {
		t.Errorf(`The user name should be "after_test", instead of %s`, resultAfterUpdate.Name)
	}
	accessToken, err := operation.Authorization("after_test", "after_test")
	if err != nil {

	}
	fmt.Println(accessToken)
}
