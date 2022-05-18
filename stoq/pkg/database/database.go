package database

import (
	"database/sql"
	"log"

	"github.com/faelp22/tcs_curso/stoq/config"
)

func NewDB(conf *config.Config) *sql.DB {
	db, err := sql.Open(conf.DB_DRIVE, conf.DB_DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
