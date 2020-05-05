package router

import (
	"github.com/gorilla/mux"
	"gitlab.devkeeper.com/authreg/server/internal/app/auth"
)

func New() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/tokens", auth.AuthMiddleware(getTokens)).Methods("GET")
	r.HandleFunc("/tokens/{id}", auth.AuthMiddleware(getToken)).Methods("GET")
	r.HandleFunc("/tokens", auth.AuthMiddleware(createToken)).Methods("POST")
	r.HandleFunc("/tokens/{id}", auth.AuthMiddleware(updateToken)).Methods("PUT")
	r.HandleFunc("/tokens/{id}", auth.AuthMiddleware(deleteToken)).Methods("DELETE")

	r.HandleFunc("/tokens/find/{key}", auth.AuthMiddleware(findTokenByKey)).Methods("GET")

	r.HandleFunc("/auth/login", auth.Login).Methods("POST")

	return r
}
