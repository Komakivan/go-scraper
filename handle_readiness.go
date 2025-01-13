package main

import "net/http"

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, map[string]string{"msg": "api healthy"})
}
