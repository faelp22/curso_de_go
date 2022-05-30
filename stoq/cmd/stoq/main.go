package main

import (
	"fmt"

	"github.com/faelp22/tcs_curso/stoq/config"
	"github.com/faelp22/tcs_curso/stoq/handler"
	"github.com/faelp22/tcs_curso/stoq/pkg/database"
	"github.com/faelp22/tcs_curso/stoq/pkg/http"
	"github.com/faelp22/tcs_curso/stoq/pkg/service"
	"github.com/faelp22/tcs_curso/stoq/web"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {

	conf := config.NewConfig("8080", config.DBConfig{
		DB_DRIVE: "sqlite3",
		// DB_HOST:  "192.168.0.100",
		// DB_PORT:  "5432",
		// DB_USER:  "root",
		// DB_PASS:  "123456",
		DB_NAME: "db.sqlite3",
	})

	dbpool := database.NewDB(conf)
	service := service.NewProdutoService(dbpool)

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.RegisterAPIHandlers(r, n, service)

	if conf.WEB_UI {
		web.RegisterUIHandlers(r, n)
	}

	if conf.Mode == config.DEVELOPER {
		r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			t, err := route.GetPathTemplate()
			if err != nil {
				return err
			}
			fmt.Println(t)
			return nil
		})
	}

	http.NewHTTPServer(r, conf)

}
