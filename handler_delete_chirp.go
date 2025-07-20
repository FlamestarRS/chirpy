package main

import (
	"net/http"

	"github.com/FlamestarRS/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, req *http.Request) {

	bearerToken, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "malformed header", err)
		return
	}

	refreshToken, _ := cfg.db.GetRefreshTokenbyID(req.Context(), bearerToken)
	if refreshToken.Token == bearerToken {
		respondWithError(w, http.StatusUnauthorized, "access token requried", err)
		return
	}

	authenticatedID, err := auth.ValidateJWT(bearerToken, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "error validating jwt", err)
		return
	}

	id, err := uuid.Parse(req.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	chirp, err := cfg.db.GetChirpByID(req.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "error: chirp not found", err)
		return
	}

	if chirp.UserID != authenticatedID {
		respondWithError(w, http.StatusForbidden, "cannot delete chirp from another user", err)
		return
	}

	err = cfg.db.DeleteChirp(req.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error deleting chirp", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
