package database

import "github.com/xalabs/refima/config"

type ServiceInterface interface {
	NewHandler(conf config.Config) HandlerInterface
}

type HandlerInterface interface {
	NewUser(name, password string) error
	GetUser(u User) (User, error)
	UpdateUser(u User) error
	Authorization(name, password string) (AccessToken, error)
	GetAccessToken(a AccessToken) (AccessToken, error)
}
