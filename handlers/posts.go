package handlers

import (
	"encoding/json"
	"myapi/models"
	"net/http"
)

func GetPosts(allPosts []*models.Post) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(allPosts)
		w.WriteHeader(200)
	}
}

func DeletePost(id int, allPosts []*models.Post) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post := r.Context().Value("post").(*models.Post)

		for i, p := range allPosts {
			if p.Id == id {
				allPosts = append(allPosts[:i], allPosts[i+1:]...)
			}
		}
		w.Write([]byte("deleted"))
	}
}
