package database

import (
	"database/sql"

	"github.com/xalabs/refima/config"
)

type SQLService struct {
	Conf config.Config
	DB   *sql.DB
}

type SQLHandler struct {
	Conf config.Config
	DB   *sql.DB
}

func NewSQLService(db *sql.DB, conf config.Config) ServiceInterface {
	return &SQLService{
		Conf: conf,
		DB:   db,
	}
}

func (s *SQLService) NewHandler(conf config.Config) HandlerInterface {
	return &SQLHandler{
		Conf: s.Conf,
		DB:   s.DB,
	}
}

func (s *SQLHandler) NewUser(name, password string) error {
	return nil
}

func (s *SQLHandler) GetUser(u User) (User, error) {
	return User{}, nil
}

func (s *SQLHandler) UpdateUser(u User) error {
	return nil
}

func (s *SQLHandler) Authorization(name, password string) (AccessToken, error) {
	return AccessToken{}, nil
}

func (s *SQLHandler) GetAccessToken(a AccessToken) (AccessToken, error) {
	return AccessToken{}, nil
}
