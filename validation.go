package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type parameters struct {
	Body string `json:"body"`
}

type errorResponse struct {
	Error string `json:"error"`	
}

type validResponse struct {
	Valid bool `json:"valid"`
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	response := errorResponse{Error: msg}
	respondWithJSON(w, code, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func handleValidation(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}
	response := validResponse{Valid: true}
	respondWithJSON(w, 200, response)

}
