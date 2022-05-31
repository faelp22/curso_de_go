package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/faelp22/tcs_curso/stoq/entity"
	"github.com/faelp22/tcs_curso/stoq/pkg/service"
	"github.com/gorilla/mux"
)

func getAllProduct(service service.ProdutoServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		vars := mux.Vars(r)

		ID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"MSG": "é necessario um ID", "codigo": 400}`))
			return
		}

		produto := service.GetByID(&ID)
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

func createProduct(service service.ProdutoServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		produto := entity.Produto{}

		err := json.NewDecoder(r.Body).Decode(&produto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"MSG": "Error to parse Product to JSON", "codigo": 500}`))
			return
		}

		last_id := service.Create(&produto)
		if last_id == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"MSG": "Error to Insert the Product", "codigo": 500}`))
			return
		}

		produto = *service.GetByID(&last_id)

		err = json.NewEncoder(w).Encode(produto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"MSG": "Error to parse Product to JSON", "codigo": 500}`))
			return
		}
	})
}

func updateProduct(service service.ProdutoServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		ID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"MSG": "é necessario um ID", "codigo": 400}`))
			return
		}

		produto := entity.Produto{}

		err = json.NewDecoder(r.Body).Decode(&produto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"MSG": "Error to parse Product to JSON", "codigo": 500}`))
			return
		}

		rows_affected := service.Update(&ID, &produto)
		if rows_affected == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"MSG": "Error to Update the Product", "codigo": 500}`))
			return
		}

		produto = *service.GetByID(&ID)

		err = json.NewEncoder(w).Encode(produto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"MSG": "Error to parse Product to JSON", "codigo": 500}`))
			return
		}
	})
}

func deleteProduct(service service.ProdutoServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		ID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"MSG": "é necessario um ID", "codigo": 400}`))
			return
		}

		rows_affected := service.Delete(&ID)
		if rows_affected == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"MSG": "Erro to Delete a Product", "codigo": 500}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"MSG": "Ok Product Deleted", "codigo": 200}`))
	})
}
