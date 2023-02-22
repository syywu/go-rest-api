package handlers

import (
	"encoding/json"
	"myapi/models"
	"net/http"
)

func GetPosts(allPosts []models.Post) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(allPosts)
	}
}
