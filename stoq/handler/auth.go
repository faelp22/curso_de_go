package handler

import (
	"encoding/json"
	"net/http"

	"github.com/faelp22/tcs_curso/stoq/entity"
)

func authLogin() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := entity.User{}

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"MSG": "Error to parse User from JSON", "codigo": 500}`))
			return
		}

		admin := entity.NewAdmin()
		token := entity.Token{
			Token: entity.USER_TOKEN,
		}

		if user.Username == admin.Username && user.Password == admin.Password {
			err = json.NewEncoder(w).Encode(token)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"MSG": "Error to parse Product to JSON", "codigo": 500}`))
				return
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"MSG": "Error to login, check username and password", "codigo": 401}`))
		}

	})
}
