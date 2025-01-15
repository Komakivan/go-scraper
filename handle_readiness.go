package main

import (
	"net/http"

	"github.com/Komakivan/go-scraper/json_res"
)

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	json_res.RespondJSON(w, http.StatusOK, map[string]string{"msg": "api healthy"})
}
