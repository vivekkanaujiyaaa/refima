package database

import (
	"os"
	"testing"
)

var (
	testDatabase = "./test.db"
	handler      HandlerInterface
	err          error
)

func init() {
	if _, err := os.Stat(testDatabase); !os.IsNotExist(err) {
		err := os.Remove(testDatabase)

		if err != nil {
			panic(err)
		}
	}

	sqlService := NewSQLService("sqlite3", testDatabase)
	handler, err = sqlService.NewHandler()
}

func TestMigrate(t *testing.T) {
	err = handler.Migrate()
	if err != nil {
		t.Errorf("Error should be nil, instead of %s", err.Error())
	}
}
