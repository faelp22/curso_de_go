package database

import (
	"database/sql"
	"log"

	"github.com/faelp22/tcs_curso/stoq/config"

	_ "github.com/mattn/go-sqlite3"
)

type dabase_pool struct {
	DB *sql.DB
}

func NewDB(conf *config.Config) (dbpool *dabase_pool) {

	if dbpool != nil {
		return dbpool
	} else {
		db, err := sql.Open("sqlite3", conf.DB_DSN)
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
