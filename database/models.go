package database

import "time"

type User struct {
	ID         string `gorm:"primary_key"`
	Name       string `gorm:"type:varchar(100);unique"`
	Password   string `gorm:"type:varchar(100)"`
	SystemUser string `gorm:"type:varchar(100)"`
}

func (User) TableName() string {
	return "users"
}

type AccessToken struct {
	ID        uint   `gorm:"primary_key"`
	UserID    string `gorm:"type:varchar(40)"`
	Token     string `gorm:"type:varchar(30);"`
	ExpiresAt time.Time
}

func (AccessToken) TableName() string {
	return "access_token"
}
