package handler

import (
	"encoding/json"
	"net/http"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request) (string, error) {
	// get string "url" from json request body and shorten it.
	defer r.Body.Close()

	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return "", nil
	}
	url, ok := data["url"]
	if !ok || url == "" {
		http.Error(w, "Missing 'url' field", http.StatusBadRequest)
		return "", nil
	}

	return url, nil
}
