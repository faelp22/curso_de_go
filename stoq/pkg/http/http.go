package http

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/faelp22/tcs_curso/stoq/config"
	"github.com/urfave/negroni"
)

func NewHTTPServer(conf *config.Config, n *negroni.Negroni) {

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + conf.SRV_PORT,
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
