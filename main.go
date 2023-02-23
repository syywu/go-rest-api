package main

import (
	"encoding/json"
	"fmt"
	"log"
	"myapi/handlers"
	"myapi/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	// allPosts := models.New()
	var allPosts = []*models.Post{
		{UserId: 100, Id: 1, Title: "hello", Body: "hello world"},
	}

	// allPosts.Add(models.Post{
	// 	UserId: 100,
	// 	Id:     1,
	// 	Title:  "hello",
	// 	Body:   "hello world",
	// })

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// get
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Root Route"))
	})

	r.Route("/posts", func(r chi.Router) {
		r.Get("/", handlers.GetPosts(allPosts))
		// post
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			req := &models.Post{}
			allPosts = append(allPosts, req)
			json.NewDecoder(r.Body).Decode(req)
			// allPosts.Add(models.Post{
			// 	UserId: req["userid"],
			// 	Id:     req["id"],
			// 	Title:  req["title"],
			// 	Body:   req["body"],
			// })

			w.Write([]byte("posted"))
			render.Status(r, http.StatusCreated)

		})
	})
	r.Route("/{postID}", func(r chi.Router) {
		// delete
		r.Delete("/posts", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("deleted"))
		})

		// patch

		// put
	})

	fmt.Print("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
