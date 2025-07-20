package main

import (
	"encoding/json"
	"net/http"

	"github.com/FlamestarRS/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUpgradeStatus(w http.ResponseWriter, req *http.Request) {
	apiKey, err := auth.GetAPIKey(req.Header)
	if apiKey != cfg.polka_key {
		respondWithError(w, http.StatusUnauthorized, "unapproved apikey", err)
		return
	}

	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	params := parameters{}
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	if params.Event != "user.upgraded" {
		respondWithError(w, http.StatusNoContent, "no upgrade event", err)
		return
	}

	id, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error parsing id", err)
		return
	}

	err = cfg.db.UpgradeChirpyRedByID(req.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error upgrading user status", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
