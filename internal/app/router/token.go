package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.devkeeper.com/authreg/server/internal/app/model"
	"gitlab.devkeeper.com/authreg/server/internal/app/store"
)

func getTokens(w http.ResponseWriter, r *http.Request) {

	tr := &store.TokenRepository{}

	tokens, err := tr.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// respondWithJson(w, http.StatusOK, tokens)

	respondWithJson(w, http.StatusOK, map[string]interface{}{
		"result": "success",
		"data":   tokens,
	})
}

func getToken(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	tr := &store.TokenRepository{}

	token, err := tr.Find(params["id"])

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]interface{}{
		"result": "success",
		"data":   token,
	})

}

func createToken(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var token *model.Token
	tr := &store.TokenRepository{}

	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := tr.Create(token)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]interface{}{
		"result": "success",
		"data":   token,
	})

}

func updateToken(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var token *model.Token
	tr := &store.TokenRepository{}

	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := tr.Update(token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]interface{}{
		"result": "success",
		"data":   token,
	})

}

func deleteToken(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	tr := &store.TokenRepository{}

	result, err := tr.Delete(params["id"])

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]interface{}{
		"result": "success",
		"data":   result,
	})
}
