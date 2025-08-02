package handler

import (
	"encoding/json"
	"net/http"
	"url-shortener/internal/service"
	"url-shortener/internal/storage"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request, s *storage.InMemoryStorage) {
	defer r.Body.Close()

	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	url, ok := data["url"]
	if !ok || url == "" {
		http.Error(w, "Missing 'url' field", http.StatusBadRequest)
		return
	}
	shortURL := service.GenerateShortUrl()
	s.Save(shortURL, url)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"short_url":"` + shortURL + `"}`))
}
