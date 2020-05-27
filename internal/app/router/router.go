package router

import (
	"github.com/gorilla/mux"
	"gitlab.devkeeper.com/authreg/server/internal/app/auth"
)

// New - router instance
func New() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/tokens", auth.MyAuthMiddleware(getTokens)).Methods("GET")
	r.HandleFunc("/tokens/{id}", auth.MyAuthMiddleware(getToken)).Methods("GET")
	r.HandleFunc("/tokens", auth.MyAuthMiddleware(createToken)).Methods("POST")
	r.HandleFunc("/tokens/{id}", auth.MyAuthMiddleware(updateToken)).Methods("PUT")
	r.HandleFunc("/tokens/{id}", auth.MyAuthMiddleware(deleteToken)).Methods("DELETE")

	r.HandleFunc("/tokens/find/{key}", findTokenByKey).Methods("GET")

	r.HandleFunc("/auth/login", auth.Login).Methods("POST")

	return r
}
