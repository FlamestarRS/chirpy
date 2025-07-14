package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

/*
func handlerValidate(w http.ResponseWriter, r *http.Request) error {
	type parameters struct {
		Body string `json:"body"`
	}
	//type returnVals struct {
	//	CleanedBody string `json:"cleaned_body"`
	//}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		return respondWithError(w, http.StatusInternalServerError, "Something went wrong", err)
	}

	if len(params.Body) > 140 {
		return respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
	}

	return nil
	//respondWithJSON(w, http.StatusOK, returnVals{CleanedBody: handlerFilter(params.Body)})
}*/

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

func respondWithError(w http.ResponseWriter, statusCode int, msg string, err error) error {
	if err != nil {
		log.Println(err)
	}
	if statusCode > 499 {
		log.Printf("internal server error: %v", statusCode)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}

	return respondWithJSON(w, statusCode, errorResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		return err
	}
	w.WriteHeader(statusCode)
	w.Write(data)
	return nil
}
