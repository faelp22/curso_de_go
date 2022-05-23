package database

import (
	"database/sql"
	"log"

	"github.com/faelp22/tcs_curso/stoq/config"

	_ "github.com/mattn/go-sqlite3"
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

	if dbpool != nil && dbpool.DB != nil {

		return dbpool

	} else {

		db, err := sql.Open(conf.DB_DRIVE, conf.DB_DSN)
		if err != nil {
			log.Fatal(err)
		}
		// defer db.Close()

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		dbpool = &dabase_pool{
			DB: db,
		}
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

func (dbp *dabase_pool) GetDB() (DB *sql.DB) {
	return dbp.DB
}
