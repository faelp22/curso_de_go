package handler

import (
	"net/http"

	"github.com/faelp22/tcs_curso/stoq/entity"
	"github.com/urfave/negroni"
)

func applicationJSON() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(w, r)
	})
}

func isAuth() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		if r.URL.Path == "/api/v1/user/login" {
			next(w, r)
			return
		}

		tmp_token := r.Header.Get("Authorization")
		if len(tmp_token) > 0 {
			if tmp_token[7:] == entity.USER_TOKEN {
				next(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"MSG": "Error token invalid", "codigo": 401}`))
	})
}
