package models

import "time"

type User struct {
	ID         string
	Name       string
	Password   string
	SystemUser string
}

func (User) TableName() string {
	return "users"
}

type AccessToken struct {
	ID        uint
	UserID    string
	Token     string
	ExpiresAt time.Time
}

func (AccessToken) TableName() string {
	return "access_token"
}
