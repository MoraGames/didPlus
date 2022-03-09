package database

import (
	"database/sql"
	"fmt"
	"github.com/MoraGames/didPlus/go/internal/config"
)

type SqlDb struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
	Db string `yaml:"db" json:"db"`
	Driver string `yaml:"driver" json:"driver"`
	Client *sql.DB
}

func NewSqlDb(cfg *config.Database) (*SqlDb, error) {
	d := &SqlDb{
		Username: cfg.Username,
		Password: cfg.Password,
		Host:     cfg.Host,
		Port:     cfg.Port,
		Db:       cfg.Db,
		Driver:   cfg.Driver,
		Client:   nil,
	}

	err := d.Connect()
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (sdb *SqlDb) Connect() (err error) {

	var url string

	if sdb.Username == "" && sdb.Password == "" {
		url = fmt.Sprintf("tcp(%s:%s)/%s", sdb.Host, sdb.Port, sdb.Db)
	} else if sdb.Username != "" && sdb.Password == "" {
		url = fmt.Sprintf("%s@tcp(%s:%s)/%s", sdb.Username, sdb.Host, sdb.Port, sdb.Db)
	} else {
		url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", sdb.Username, sdb.Password, sdb.Host, sdb.Port, sdb.Db)
	}

	sdb.Client, err = sql.Open(sdb.Driver, url)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

func (sdb *SqlDb) Close() (err error) {
	return sdb.Client.Close()
}
