package database

import (
	"database/sql"
	"fmt"

	"github.com/faelp22/tcs_curso/stoq/config"
)

type DatabaseInterface interface {
	GetDB() (DB *sql.DB)
	Close() error
}

type dabase_pool struct {
	DB *sql.DB
}

var dbpool = &dabase_pool{}

func NewDB(conf *config.Config) *dabase_pool {

	switch conf.DBConfig.DB_DRIVE {
	case "sqlite3":
		conf.DBConfig.DB_DSN = fmt.Sprintf(conf.DB_NAME)
		dbpool = SQLiteConn(conf)
	case "postgres":
		conf.DBConfig.DB_DSN = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.DB_HOST, conf.DB_PORT, conf.DB_USER, conf.DB_PASS, conf.DB_NAME)
		dbpool = PGConn(conf)
	default:
		panic("Drive não implementado")
	}

	return dbpool
}

func (d *dabase_pool) Close() error {

	err := d.DB.Close()
	if err != nil {
		return err
	}

	dbpool = &dabase_pool{}

	return err
}

func (d *dabase_pool) GetDB() (DB *sql.DB) {
	return d.DB
}
