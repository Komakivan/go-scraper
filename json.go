package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "something went wrong"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func ResponseError(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		log.Println("Responding with 5XX error message: ", msg)
	}
	type ErrorResponse struct {
		Error string `json:"error"`
	}
	RespondJSON(w, status, ErrorResponse{
		Error: msg,
	})
}
