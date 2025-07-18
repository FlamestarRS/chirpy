package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FlamestarRS/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Email      string        `json:"email"`
		Password   string        `json:"password"`
		Expiration time.Duration `json:"expires_in_seconds"`
	}
	params := parameters{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong", err)
		return
	}
	if params.Expiration <= 0 || params.Expiration > 3600 {
		params.Expiration = time.Second * 3600
	}
	user, err := cfg.db.GetUserByEmail(req.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "incorrect email or password", nil)
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "incorrect email or password", nil)
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.secret, params.Expiration)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "token expired", nil)
	}

	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Token:     token,
	})
}
