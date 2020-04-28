package router

import (
	"net/http"

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
