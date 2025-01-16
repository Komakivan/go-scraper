package main

import (
	"time"

	"github.com/Komakivan/go-scraper/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Api_key   string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedsFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func sanitizeUser(user database.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Api_key:   user.ApiKey,
	}
}

func sanitizeFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		Name:      feed.Name,
		Url:       feed.Url,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		UserID:    feed.UserID,
	}
}

func sanitizeFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbfeed := range dbFeeds {
		feeds = append(feeds, sanitizeFeed(dbfeed))
	}
	return feeds
}

func sanitizeFeedFollow(dbFeed database.FeedsFollow) FeedsFollow {
	return FeedsFollow{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		UserID:    dbFeed.UserID,
		FeedID:    dbFeed.FeedID,
	}
}

func sanitizeFeedFollows(dbFeedFollows []database.FeedsFollow) []FeedsFollow {
	feedsFollow := []FeedsFollow{}
	for _, feedFollow := range dbFeedFollows {
		feedsFollow = append(feedsFollow, sanitizeFeedFollow(feedFollow))
	}
	return feedsFollow
}
