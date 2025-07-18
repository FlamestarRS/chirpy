package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/FlamestarRS/chirpy/internal/auth"
	"github.com/FlamestarRS/chirpy/internal/database"
	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerChirp(w http.ResponseWriter, req *http.Request) {
	type requestParams struct {
		Body   string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}

	params := requestParams{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	bearerToken, err := auth.GetBearerToken(req.Header)

	if err != nil {
		fmt.Println("error getting bearer token")
		return
	}
	authenticatedID, err := auth.ValidateJWT(bearerToken, cfg.secret)

	if err != nil {
		fmt.Println("error validating jwt")
		return
	}
	err = validateChirp(w, params.Body)
	if err != nil {
		return
	}

	chirp, err := cfg.db.CreateChirp(req.Context(), database.CreateChirpParams{
		Body:   params.Body,
		UserID: authenticatedID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating chirp", err)
		return
	}

	respondWithJSON(w, http.StatusCreated,
		Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      handlerFilter(chirp.Body),
			UserID:    chirp.UserID,
		})

}

func validateChirp(w http.ResponseWriter, text string) error {
	if len(text) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return fmt.Errorf("chirp cannot exceed 140 characters")
	}
	return nil
}

func handlerFilter(text string) string {
	wordsToFilter := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}

	words := strings.Split(text, " ")

	for i, word := range words {
		lowered := strings.ToLower(word)
		if _, ok := wordsToFilter[lowered]; ok {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
}
