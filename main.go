package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Komakivan/go-scraper/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port must be set")
	}
	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Fatal("database string must be set")
	}

	conn, err := sql.Open("postgres", dbString)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	apiConfig := &ApiConfig{
		DB: database.New(conn),
	}

	// create a router
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// --------user routes -------//
	v1router := chi.NewRouter()
	v1router.Get("/health", HandleReadiness)
	v1router.Get("/error", HandlerError)
	v1router.Post("/users", apiConfig.HandlerCreateUser)
	v1router.Get("/user", apiConfig.authMiddleware(apiConfig.HandlerGetUserByApiKey))

	// -------feed routes ----------//
	v1router.Post("/feeds", apiConfig.authMiddleware(apiConfig.handlerCreateFeed)) // authenticated endpoint
	v1router.Get("/feeds", apiConfig.handlerGetFeeds)
	v1router.Post("/feeds/follow", apiConfig.authMiddleware(apiConfig.HandlerFollowFeed))
	v1router.Get("/feeds/follows", apiConfig.authMiddleware(apiConfig.HandlerGetFeedFollows))
	v1router.Delete("/feeds/follow/{feedFollowId}", apiConfig.authMiddleware(apiConfig.HandlerDeleteFeedFollow))

	router.Mount("/v1", v1router)

	server := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	log.Println("Server starting on port: ", portString)
	err = server.ListenAndServe()
	if err != nil {
		panic("failed to start server")
	}
}
