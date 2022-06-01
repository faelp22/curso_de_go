package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/faelp22/tcs_curso/stoq/config"
	"github.com/faelp22/tcs_curso/stoq/handler"
	"github.com/faelp22/tcs_curso/stoq/pkg/database"
	"github.com/faelp22/tcs_curso/stoq/pkg/http"
	"github.com/faelp22/tcs_curso/stoq/pkg/service"
	"github.com/faelp22/tcs_curso/stoq/webui"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {

	conf := config.NewConfig(config.DBConfig{
		DB_DRIVE: "sqlite3",
		// DB_HOST:  "192.168.0.100",
		// DB_PORT:  "5432",
		// DB_USER:  "root",
		// DB_PASS:  "123456",
		DB_NAME: "db.sqlite3",
	})

	conf.WEB_UI = true

	dbpool := database.NewDB(conf)
	service := service.NewProdutoService(dbpool)

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	if conf.WEB_UI {
		webui.RegisterUIHandlers(r, n)
	}

	handler.RegisterAPIHandlers(r, n, service)

	if conf.Mode == config.DEVELOPER {
		r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			t, err := route.GetPathTemplate()
			if err != nil {
				return err
			}
			fmt.Println(strings.Replace(t, `{_dummy:.*}`, "*", 1))
			return nil
		})
	}

	srv := http.NewHTTPServer(r, conf)

	done := make(chan bool)
	go srv.ListenAndServe()
	log.Printf("Server Run on Port: %v, Mode: %v, WEBUI: %v", conf.SRV_PORT, conf.Mode, conf.WEB_UI)
	<-done

}
