package main

import (
	"github.com/faelp22/tcs_curso/stoq/config"
	"github.com/faelp22/tcs_curso/stoq/handler"
	"github.com/faelp22/tcs_curso/stoq/pkg/database"
	"github.com/faelp22/tcs_curso/stoq/pkg/http"
	"github.com/faelp22/tcs_curso/stoq/pkg/service"
	"github.com/urfave/negroni"
)

func main() {

	config := config.NewConfig("8080", config.DBConfig{
		DB_DRIVE: "sqlite3",
		// DB_HOST:  "192.168.0.100",
		// DB_PORT:  "5432",
		// DB_USER:  "root",
		// DB_PASS:  "123456",
		DB_NAME: "db.sqlite3",
	}, false)

	dbpool := database.NewDB(config)
	service := service.NewProdutoService(dbpool)

	n := negroni.Classic()
	handler.NewRouter(n, service)
	http.NewHTTPServer(config, n)
}
