package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, req *http.Request) {
	allChirps, err := cfg.db.GetAllChirps(req.Context())
	if err != nil {
		fmt.Printf("error getting chirps: %v", err)
		return
	}
	formatted := []Chirp{} // required to correctly format json tags eg CreatedAt -> created_at
	for _, chirp := range allChirps {
		formatted = append(formatted, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.ID,
		})
	}

	respondWithJSON(w, http.StatusOK, formatted)
}

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, req *http.Request) {
	id, err := uuid.Parse(req.PathValue("chirpID"))
	if err != nil {
		fmt.Printf("error parsing id: %v", err)
		return
	}
	chirp, err := cfg.db.GetChirpByID(req.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "error: chirp not found", nil)
		return
	}

	type response struct {
		Chirp
	}

	respondWithJSON(w, http.StatusOK, response{Chirp: Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}})
}
