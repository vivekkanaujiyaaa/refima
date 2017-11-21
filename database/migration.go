package database

import "github.com/xalabs/refima/config"

func Migrate(conf config.Config) error {
	db, err := connection(conf)
	if err != nil {
		return err
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserUID{})
	db.AutoMigrate(&UserGID{})
	db.AutoMigrate(&AccessToken{})
	return nil
}
