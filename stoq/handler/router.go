package handler

import (
	"github.com/faelp22/tcs_curso/stoq/pkg/service"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func RegisterAPIHandlers(r *mux.Router, n *negroni.Negroni, service service.ProdutoServiceInterface) {

	api := r.PathPrefix("/api/v1").Subrouter()

	n.Use(applicationJSON())
	n.Use(isAuth())

	api.Handle("/user/login", n.With(
		negroni.Wrap(authLogin()),
	)).Methods("POST", "OPTIONS")

	api.Handle("/products", n.With(
		negroni.Wrap(getAllProduct(service)),
	)).Methods("GET", "OPTIONS")

	api.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	api.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	api.Handle("/product/{id}", n.With(
		negroni.Wrap(updateProduct(service)),
	)).Methods("PUT", "OPTIONS")

	api.Handle("/product/{id}", n.With(
		negroni.Wrap(deleteProduct(service)),
	)).Methods("DELETE", "OPTIONS")

}
