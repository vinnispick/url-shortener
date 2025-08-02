package server

import (
	"net/http"
	"url-shortener/internal/handler"
	"url-shortener/internal/storage"
)

func NewRouter(s *storage.InMemoryStorage) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.RedirectHandler(w, r, s)
		case http.MethodPost:
			handler.ShortenHandler(w, r, s)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return router
}
