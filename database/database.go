package database

import (
	"github.com/PumpkinSeed/refima/config"
	"github.com/jinzhu/gorm"
)

type Operation struct {
	Conf config.Config
	DB   *gorm.DB
}

func NewOperationHandler(conf config.Config) (Operation, error) {
	db, err := connection(conf)
	if err != nil {
		return Operation{}, err
	}
	return Operation{
		Conf: conf,
		DB:   db,
	}, nil
}
