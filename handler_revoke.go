package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, req *http.Request) {
	TOKEN_STRING := strings.Trim(strings.TrimPrefix(req.Header.Get("Authorization"), "Bearer "), " ")
	if len(TOKEN_STRING) <= 0 {
		fmt.Println("no token", TOKEN_STRING)
		return

	}

	refreshToken, err := cfg.db.GetRefreshTokenbyID(req.Context(), TOKEN_STRING)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "refresh token does not exist", err)
		return
	}

	cfg.db.RevokeRefreshToken(req.Context(), refreshToken.Token)
	respondWithJSON(w, http.StatusNoContent, nil)
}
