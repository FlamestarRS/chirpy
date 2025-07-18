package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/FlamestarRS/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, req *http.Request) {
	TOKEN_STRING := strings.Trim(strings.TrimPrefix(req.Header.Get("Authorization"), "Bearer "), " ")
	if len(TOKEN_STRING) <= 0 {
		fmt.Println("no token", TOKEN_STRING)
		return

	}

	refreshToken, err := cfg.db.GetRefreshTokenbyID(req.Context(), TOKEN_STRING)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "refresh token does not exist", nil)
		return
	}

	if refreshToken.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "refresh token revoked", nil)
		return
	}

	durationUntil := time.Until(refreshToken.ExpiresAt)
	if durationUntil <= 0 {
		respondWithError(w, http.StatusUnauthorized, "refresh token expired", nil)
		return
	}

	type response struct {
		Token string `json:"token"`
	}
	newJwt, err := auth.MakeJWT(refreshToken.UserID, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating new jwt", nil)
	}
	respondWithJSON(w, http.StatusOK, response{
		Token: newJwt,
	})

}
