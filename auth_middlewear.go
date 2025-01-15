package main

import (
	"log"
	"net/http"

	"github.com/Komakivan/go-scraper/internal/auth"
	"github.com/Komakivan/go-scraper/internal/database"
	"github.com/Komakivan/go-scraper/json_res"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) authMiddleware(next authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			json_res.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			log.Println(err)
			json_res.ResponseError(w, http.StatusNotFound, "user does not exist")
			return
		}

		next(w, r, user)
	}
}
