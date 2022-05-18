package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(srv string) {

	r := mux.NewRouter()
	r.HandleFunc("/products/{key}", ProductHandler)
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	http.Handle("/", r)
}
