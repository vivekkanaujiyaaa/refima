package database

import "github.com/jinzhu/gorm"

func connection() {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
}
