package webui

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//go:embed dist/*
var Assets embed.FS

type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

func AssetHandler(prefix, root string) http.Handler {

	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join(root, name)

		f, err := Assets.Open(assetPath)
		if os.IsNotExist(err) {
			return Assets.Open("dist/spa/index.html")
		}

		return f, err
	})

	return http.StripPrefix(prefix, http.FileServer(http.FS(handler)))
}

func RegisterUIHandlers(r *mux.Router, n *negroni.Negroni) {

	r.PathPrefix("/webui/{_dummy:.*}").Handler(n.With(
		negroni.Wrap(AssetHandler("/webui", "dist/spa/")),
	)).Methods("GET", "OPTIONS")

}
