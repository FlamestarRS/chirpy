package main

import (
	"fmt"
	"net/http"
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
