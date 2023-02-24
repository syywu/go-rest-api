package handlers

import (
	"myapi/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// func GetPosts(allPosts []*models.Post) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		json.NewEncoder(w).Encode(allPosts)
// 		w.WriteHeader(200)
// 	}
// }

// func DeletePost(id int, allPosts []*models.Post) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		post := r.Context().Value("post").(*models.Post)

// 		for i, p := range allPosts {
// 			if p.Id == id {
// 				allPosts = append(allPosts[:i], allPosts[i+1:]...)
// 			}
// 		}
// 		w.Write([]byte("deleted"))
// 	}
// }

var postIDkey = "postID"

func posts(r chi.Router) {
	r.Get("/", getAllPosts)
	r.Post("/", createPost)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{}
	if err := render.Bind(r, post); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := DBInstance.AddPost(post); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, post); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
