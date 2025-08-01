package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"url-shortener/internal/handler"
	"url-shortener/internal/service"
	"url-shortener/internal/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if host == "" || port == "" {
		panic("Environment variables HOST and PORT must be set")
	}

	storage := storage.NewInMemoryStorage()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			path := strings.TrimPrefix(r.URL.Path, "/")
			fmt.Println("storage data:", storage)
			log.Println("Received request for path:", path)
			if path == "" {
				log.Println("path is empty, redirecting to home")
				http.NotFound(w, r)
				return
			}
			originalURL, exists := storage.Get(path)
			if !exists {
				http.NotFound(w, r)
			}
			http.Redirect(w, r, originalURL, http.StatusFound)
			return
		case http.MethodPost:
			url, err := handler.ShortenHandler(w, r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			shortURL := service.GenerateShortUrl()
			storage.Save(shortURL, url)
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"short_url":"` + shortURL + `"}`))
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	http.ListenAndServe(":"+port, nil)
}
