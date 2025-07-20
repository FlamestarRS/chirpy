package main

import (
	"encoding/json"
	"net/http"

	"github.com/FlamestarRS/chirpy/internal/auth"
	"github.com/FlamestarRS/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUpdatePassword(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params := parameters{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

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

	newPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
		return
	}

	err = cfg.db.UpdatePassword(req.Context(), database.UpdatePasswordParams{
		Email:          params.Email,
		HashedPassword: newPassword,
		ID:             authenticatedID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error updating password", err)
		return
	}

	respondWithJSON(w, http.StatusOK, User{
		Email: params.Email,
	})
}
