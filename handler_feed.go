package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Komakivan/go-scraper/internal/database"
	"github.com/Komakivan/go-scraper/json_res"
	"github.com/google/uuid"
)

// Authenticated endpoint
func (apiCfg *ApiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params := parameters{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		json_res.ResponseError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	defer r.Body.Close()

	// create the feed
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Url:       params.Url,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
	})

	if err != nil {
		log.Println(err)
		json_res.ResponseError(w, http.StatusInternalServerError, "Could not create feed")
		return
	}

	json_res.RespondJSON(w, http.StatusCreated, sanitizeFeed(feed))

}

func (apiCfg *ApiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		log.Println()
		json_res.ResponseError(w, http.StatusInternalServerError, "failed to get feeds")
		return
	}

	json_res.RespondJSON(w, http.StatusOK, sanitizeFeeds(feeds))
}
