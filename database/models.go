package database

import "time"

type User struct {
	ID       string `gorm:"primary_key"`
	Name     string `gorm:"type:varchar(100);unique"`
	Password string `gorm:"type:varchar(100)"`
}

func (User) TableName() string {
	return "users"
}

type UserUID struct {
	ID     uint   `gorm:"primary_key"`
	UserID string `gorm:"type:varchar(40)"`
	UID    string `gorm:"type:varchar(30);"`
}

func (UserUID) TableName() string {
	return "user_uid"
}

type UserGID struct {
	ID     uint   `gorm:"primary_key"`
	UserID string `gorm:"type:varchar(40)"`
	GID    string `gorm:"type:varchar(30);"`
}

func (UserGID) TableName() string {
	return "user_gid"
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
