package handler

import (
	"net/http"

	"github.com/faelp22/tcs_curso/stoq/pkg/service"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func NewRouter(n *negroni.Negroni, service service.ProdutoServiceInterface) {
	mux := mux.NewRouter()
	mux.Handle("/products/", getAllProduct(service)).Methods("GET")
	mux.Handle("/product/{id}", getProduct(service)).Methods("GET")
	mux.Handle("/product", createProduct(service)).Methods("POST")
	mux.Handle("/product/{id}", updateProduct(service)).Methods("PUT")
	mux.Handle("/product/{id}", deleteProduct(service)).Methods("DELETE")
	n.UseHandler(mux)
	http.Handle("/", n)
}
