package database

import (
	"database/sql"
	"log"

	"github.com/faelp22/tcs_curso/stoq/config"

	_ "github.com/mattn/go-sqlite3"
)

func SQLiteConn(conf *config.Config) *dabase_pool {

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
