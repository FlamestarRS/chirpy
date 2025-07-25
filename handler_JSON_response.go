package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if statusCode > 499 {
		log.Printf("internal server error: %v", statusCode)
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
		w.WriteHeader(500)
		log.Print(err)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(data)
}
