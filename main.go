package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	portString := os.Getenv("PORT")

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

	// handler routes
	v1router := chi.NewRouter()
	v1router.Get("/health", HandleReadiness)
	v1router.Get("/error", HandlerError)

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
