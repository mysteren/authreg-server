package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/tokens", getTokens).Methods("GET")
	// r.HandleFunc("/tokens/{id}", getToken).Methods("GET")
	// r.HandleFunc("/tokens", createtoken).Methods("POST")
	// r.HandleFunc("/tokens/{id}", updateToken).Methods("PUT")
	// r.HandleFunc("/tokens/{id}", deleteToken).Methods("DELETE")

	return r
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
