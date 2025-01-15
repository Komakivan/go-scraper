package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Komakivan/go-scraper/internal/database"
	"github.com/Komakivan/go-scraper/json_res"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		json_res.ResponseError(w, http.StatusBadRequest, fmt.Sprintf("error parsing json: %v", err))
		return
	}
	defer r.Body.Close()

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		log.Println(err)
		json_res.ResponseError(w, http.StatusInternalServerError, "failed to create user")
		return
	}

	json_res.RespondJSON(w, http.StatusCreated, sanitizeUser(user))
}

func (apiCfg *ApiConfig) HandlerGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {

	json_res.RespondJSON(w, http.StatusOK, sanitizeUser(user))
}
