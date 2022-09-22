package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/faelp22/tcs_curso/stoq/config"
	"github.com/faelp22/tcs_curso/stoq/handler"
	"github.com/faelp22/tcs_curso/stoq/pkg/database"
	lhttp "github.com/faelp22/tcs_curso/stoq/pkg/http"
	"github.com/faelp22/tcs_curso/stoq/pkg/service"
	"github.com/faelp22/tcs_curso/stoq/webui"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {

	default_conf := &config.Config{}

	if file_config := os.Getenv("STOQ_CONFIG"); file_config != "" {
		file, _ := os.ReadFile(file_config)
		_ = json.Unmarshal(file, &default_conf)
	}

	conf := config.NewConfig(default_conf)

	dbpool := database.NewDB(conf)
	service := service.NewProdutoService(dbpool)

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	r.HandleFunc("/", redirect)

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

	srv := lhttp.NewHTTPServer(r, conf)

	done := make(chan bool)
	go srv.ListenAndServe()
	log.Printf("Server Run on Port: %v, Mode: %v, DB-Driver: %v, WEBUI: %v", conf.SRV_PORT, conf.Mode, conf.DBConfig.DB_DRIVE, conf.WEB_UI)
	open(fmt.Sprintf("http://localhost:%v", conf.SRV_PORT), conf)
	<-done

}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/webui/", http.StatusMovedPermanently)
}

func open(url string, conf *config.Config) error {
	var cmd string
	var args []string

	if !conf.OpenBrowser {
		return nil
	}

	switch runtime.GOOS {
	case "windows": // For Windows
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin": // Mac OS
		cmd = "open"
	default: // For "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
