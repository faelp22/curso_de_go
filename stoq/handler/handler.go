package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/faelp22/tcs_curso/stoq/pkg/service"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func NewRouter(n *negroni.Negroni, service service.ProdutoServiceInterface) {
	mux := mux.NewRouter()
	mux.Handle("/products/", getAllProduct(service))
	mux.Handle("/product/{id}", getProduct(service))
	n.UseHandler(mux)
	http.Handle("/", n)
}

func getAllProduct(service service.ProdutoServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		all := service.GetAll()
		err := json.NewEncoder(w).Encode(all)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(all.String()))
			return
		}
	})
}

func getProduct(service service.ProdutoServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		vars := mux.Vars(r)

		ID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"MSG": "é necessario um ID", "codigo": 400}`))
			return
		}

		produto := service.GetByID(ID)
		if produto.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"MSG": "Product not found", "codigo": 404}`))
			return
		}

		err = json.NewEncoder(w).Encode(produto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"MSG": "Error to parse Product to JSON", "codigo": 500}`))
			return
		}
	})
}
