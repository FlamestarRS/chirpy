package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUpgradeStatus(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	params := parameters{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	if params.Event != "user.upgrade" {
		respondWithError(w, http.StatusNoContent, "no upgrade event", nil)
	}

	id, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error parsing id", nil)
	}

	err = cfg.db.UpgradeChirpyRedByID(req.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error upgrading user status", nil)
	}

	w.WriteHeader(http.StatusNoContent)
}
