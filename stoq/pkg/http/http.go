package http

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/faelp22/tcs_curso/stoq/config"
	"github.com/gorilla/mux"
)

func NewHTTPServer(r *mux.Router, conf *config.Config) *http.Server {

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + conf.SRV_PORT,
		Handler:      r,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}

	return srv
}
