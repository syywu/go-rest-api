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
	db.CreateTable()
	// r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	// r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root route"))
	})

	r.Get("/posts", handlers.GetAllPosts)
	r.Post("/posts", handlers.AddPost)
	r.Get("/posts/{id}", handlers.GetPostByID)
	r.Delete("/posts/{id}", handlers.DeletePost)
	r.Put("/posts/{id}", handlers.UpdatePost)

	fmt.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
