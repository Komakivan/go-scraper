package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Komakivan/go-scraper/internal/database"
	"github.com/Komakivan/go-scraper/json_res"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) HandlerFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := parameters{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		json_res.ResponseError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	// follow feed
	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		log.Println(err)
		json_res.ResponseError(w, http.StatusInternalServerError, "Failed to follow feed")
		return
	}

	json_res.RespondJSON(w, http.StatusCreated, sanitizeFeedFollow(feed_follow))
}

// authenticated endpoint
func (apiCfg *ApiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		log.Println(err)
		json_res.ResponseError(w, http.StatusInternalServerError, "failed to get feeds follow")
		return
	}

	json_res.RespondJSON(w, http.StatusOK, sanitizeFeedFollows(feedFollows))
}

func (apiCfg *ApiConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowIDStr := chi.URLParam(r, "feedFollowId")

	feedFollowUUID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		json_res.ResponseError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	// TODO: check first if that item exists in the database

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowUUID,
		UserID: user.ID,
	})

	if err != nil {
		log.Println(err)
		json_res.RespondJSON(w, http.StatusInternalServerError, "failed to unfollow feed")
		return
	}

	json_res.RespondJSON(w, http.StatusOK, map[string]string{"message": "feed successfully unfollowed"})
}
