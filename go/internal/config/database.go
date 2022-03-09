package config

import (
	"errors"
	"github.com/MoraGames/didPlus/go/internal/utils"
)

type Database struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
	Db string `yaml:"db" json:"db"`
	Driver string `yaml:"driver" json:"driver"`
}

func (db *Database) CheckDatabase() error {

	if err := db.CheckHost(); err != nil {
		return err
	}

	if err := db.CheckPort(); err != nil {
		return err
	}

	if err := db.CheckDBName(); err != nil {
		return err
	}

	return nil
}

func (db *Database) CheckHost() error {

	if db.Host == "" {
		return errors.New("Host not setted.")
	}

	return nil
}

func (db *Database) CheckPort() error {

	if err := utils.PortValid(db.Port); err != nil {
		return errors.New("Port not valid.")
	}

	return nil
}

func (db *Database) CheckDBName() error {

	if db.Db == "" {
		return errors.New("Database name not valid.")
	}

	return nil
}