package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/xalabs/refima/config"
)

func connection(conf config.Config) (*gorm.DB, error) {
	d := conf.Database

	db, err := gorm.Open(d.Dialect, d.Connection)
	db.LogMode(conf.DebugMode)
	return db, err
}
