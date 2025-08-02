package handler

import (
	"net/http"
	"strings"
	"url-shortener/internal/storage"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request, s *storage.InMemoryStorage) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" {
		http.NotFound(w, r)
		return
	}
	originalURL, exists := s.Get(path)
	if !exists {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
	w.WriteHeader(http.StatusFound)
	w.Header().Set("Location", originalURL)
	http.Redirect(w, r, originalURL, http.StatusFound)
}
