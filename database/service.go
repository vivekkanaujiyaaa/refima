package database

import "github.com/xalabs/refima/database/models"

type ServiceInterface interface {
	NewHandler(connection string) HandlerInterface
}

type HandlerInterface interface {
	NewUser(name, password, systemUser string) error
	GetUser(u models.User) (*models.User, error)
	UpdateUser(u models.User) error
	Authorization(name, password string) (*models.AccessToken, error)
	GetAccessToken(a models.AccessToken) (*models.AccessToken, error)
	Migrate() error
}
