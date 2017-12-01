package database

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"io"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xalabs/refima/database/models"
)

type SQLService struct {
	Engine     string
	Connection string
}

type SQLHandler struct {
	DB *sql.DB
}

func NewSQLService(engine, connection string) ServiceInterface {
	return &SQLService{
		Engine:     engine,
		Connection: connection,
	}
}

func (s *SQLService) NewHandler() (HandlerInterface, error) {
	db, err := sql.Open(s.Engine, s.Connection)
	if err != nil {
		return nil, err
	}

	return &SQLHandler{
		DB: db,
	}, nil
}

func (s *SQLHandler) NewUser(name, password, systemUser string) error {
	stmt, err := s.DB.Prepare("INSERT INTO users(id, name, password, system_user) VALUES(?,?,?,?)")
	// @todo add support for Postgres and Oracle
	if err != nil {
		return err
	}
	uuid, err := s.newUUID()
	if err != nil {
		return err
	}
	res, err := stmt.Exec(uuid, name, password, systemUser)
	if err != nil {
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowCnt != 1 {
		return errors.New("insert new user failed")
	}
	return nil
}

func (s *SQLHandler) GetUser(u models.User) (*models.User, error) {
	rows, err := s.DB.Query("SELECT id, name, password, system_user FROM users WHERE id = ?", u.ID)
	// @todo get user by id OR name OR system_user
	// @todo add support for Postgres and Oracle
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var res models.User
	for rows.Next() {
		err := rows.Scan(&res.ID, &res.Name, &res.Password, &res.SystemUser)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *SQLHandler) UpdateUser(u models.User) error {
	return nil
}

func (s *SQLHandler) Authorization(name, password string) (*models.AccessToken, error) {
	return &models.AccessToken{}, nil
}

func (s *SQLHandler) GetAccessToken(a models.AccessToken) (*models.AccessToken, error) {
	return &models.AccessToken{}, nil
}

func (s *SQLHandler) Migrate() error {
	return nil
}

func (s *SQLHandler) newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
