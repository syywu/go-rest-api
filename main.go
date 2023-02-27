package main

import (
	"fmt"
	"log"
	"myapi/db"
	"myapi/handlers"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()
	// r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	// r.Use(middleware.URLFormat)
	// r.Use(middleware.Recoverer)
	// r.Use(render.SetContentType(render.ContentTypeJSON))

	db.CreateTable()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root route"))
	})

	r.Route("/posts", func(r chi.Router) {
		r.Get("/", handlers.GetAllPosts)
		r.Post("/", handlers.AddPost)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetPostByID)
			r.Delete("/", handlers.DeletePost)
			r.Put("/", handlers.UpdatePost)
		})
	})

	fmt.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
