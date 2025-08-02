package main

import (
	"log"
	"net/http"
	"os"
	"url-shortener/internal/app/server"
	"url-shortener/internal/config"
	"url-shortener/internal/logger"
	"url-shortener/internal/middleware"
	"url-shortener/internal/storage"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	config := config.NewConfig(os.Getenv("HOST"), os.Getenv("PORT"))

	if config.Host == "" || config.Port == "" {
		panic("Invalid configuration: HOST and PORT must be set")
	}
	storage := storage.NewInMemoryStorage()
	logger := logger.NewStdLogger()

	router := server.NewRouter(storage)

	loggedRouter := middleware.LoggingMiddleware(logger)(router)

	log.Printf("Server listening on %s", config.ListenPort())

	if err := http.ListenAndServe(config.ListenPort(), loggedRouter); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
