package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handlerValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		Valid bool `json:"valid"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	respondWithJSON(w, http.StatusOK, returnVals{Valid: true})
}

func respondWithError(w http.ResponseWriter, statusCode int, msg string, err error) {
	if err != nil {
		fmt.Println(err)
	}
	if statusCode > 499 {
		fmt.Printf("internal server error: %v\n", statusCode)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, statusCode, errorResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(data)
}
