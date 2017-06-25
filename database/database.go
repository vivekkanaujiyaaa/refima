package database

import (
	"github.com/PumpkinSeed/refima/config"
	"github.com/PumpkinSeed/tuid"
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

func (o *Operation) NewUser(name, password string) error {
	g := tuid.NewGenerator(5, true, false)
	id, err := g.New()
	if err != nil {
		return err
	}
	ePassword, err := HashPasswordForClient(password)
	if err != nil {
		return err
	}
	user := User{
		ID:       id,
		Name:     "test",
		Password: string(ePassword),
	}
	o.DB.Create(&user)
	return nil
}

func (o *Operation) GetUser(u User) User {
	user := User{}
	o.DB.Where(&u).First(&user)
	return user
}

func (o *Operation) NewUserUID(userID, uid string) {
	u := UserUID{
		UserID: userID,
		UID:    uid,
	}
	o.DB.Create(&u)
}

func (o *Operation) NewUserGID(userID, gid string) {
	u := UserGID{
		UserID: userID,
		GID:    gid,
	}
	o.DB.Create(&u)
}

func (o *Operation) NewAccessToken(a AccessToken) {
	o.DB.Create(&a)
}
